package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"sync"

	filetransfer "s3-example/api/gen/go"
	"s3-example/internal/clients"
	"s3-example/internal/config"
	"s3-example/internal/storage"
)

type FileHandler struct {
	cfg               *config.TransferServiceConfig
	grpcClientManager *clients.GrpcClientManager
	dbManager         *storage.Manager
}

func NewFileHandler(cfg *config.TransferServiceConfig, grpcClientManager *clients.GrpcClientManager, dbManager *storage.Manager) *FileHandler {
	return &FileHandler{
		cfg:               cfg,
		grpcClientManager: grpcClientManager,
		dbManager:         dbManager,
	}
}

func calculateChunkHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func (h *FileHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, h.cfg.MaxUploadSize)

	err := r.ParseMultipartForm(h.cfg.MaxUploadSize)
	if err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	grpcClients := h.grpcClientManager.GetClients()
	clientCount := int32(len(grpcClients))
	if clientCount == 0 {
		http.Error(w, "No available gRPC connections", http.StatusInternalServerError)
		return
	}

	bufferSize := h.cfg.ChunkSize
	buffer := make([]byte, bufferSize)
	chunkNumber := int32(0)

	totalChunks := int32(0)
	totalSize := int64(0)
	contentLength := r.ContentLength

	if contentLength > 0 {
		totalChunks = int32((contentLength + int64(bufferSize) - 1) / int64(bufferSize))
		fmt.Printf("Started uploading file '%s' with size %d bytes\n", handler.Filename, contentLength)
	}

	fileID, err := h.dbManager.CreateFileMetadata(handler.Filename, totalChunks, totalSize)
	if err != nil {
		http.Error(w, "Error adding file to database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	chunksMap := make(map[string][]*filetransfer.FileChunk)

	serviceNames := h.grpcClientManager.GetClientNames()
	if len(serviceNames) == 0 {
		http.Error(w, "No available gRPC service names", http.StatusInternalServerError)
		return
	}

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if bytesRead == 0 {
			break
		}

		chunkData := make([]byte, bytesRead)
		copy(chunkData, buffer[:bytesRead])

		totalSize += int64(bytesRead)

		serviceIndex := chunkNumber % clientCount
		serviceName := serviceNames[serviceIndex]

		chunkHash := calculateChunkHash(chunkData)

		chunk := &filetransfer.FileChunk{
			Filename:    handler.Filename,
			Chunk:       chunkData,
			ChunkNumber: chunkNumber,
			TotalChunks: totalChunks,
			ChunkHash:   chunkHash,
			ServiceName: serviceName,
		}
		chunksMap[serviceName] = append(chunksMap[serviceName], chunk)

		err = h.dbManager.SaveChunkMetadata(storage.ChunkMetadata{
			FileID:      fileID,
			ChunkNumber: chunkNumber,
			ServiceName: serviceName,
			ChunkSize:   int64(bytesRead),
			ChunkHash:   chunkHash,
		})
		if err != nil {
			http.Error(w, "Error saving chunk metadata: "+err.Error(), http.StatusInternalServerError)
			return
		}

		chunkNumber++
		if totalChunks > 0 {
			percent := float64(chunkNumber) / float64(totalChunks) * 100
			fmt.Printf("Uploaded %d/%d chunks (%.2f%%)\n", chunkNumber, totalChunks, percent)
		} else {
			fmt.Printf("Uploaded %d chunks\n", chunkNumber)
		}
	}

	err = h.dbManager.UpdateFileMetadata(fileID, chunkNumber, totalSize)
	if err != nil {
		http.Error(w, "Error updating file information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Upload completed. Total chunks: %d\n", totalChunks)

	var wg sync.WaitGroup
	errCh := make(chan error, len(chunksMap))

	for serviceName, chunks := range chunksMap {
		wg.Add(1)
		go func(serviceName string, chunks []*filetransfer.FileChunk) {
			defer wg.Done()
			client := h.grpcClientManager.GetClientByName(serviceName)
			if client == nil {
				errCh <- fmt.Errorf("gRPC client not found for service: %s", serviceName)
				return
			}
			err := h.grpcClientManager.SendChunks(client, chunks)
			if err != nil {
				errCh <- fmt.Errorf("Error sending chunks via gRPC: %v", err)
				return
			}
		}(serviceName, chunks)
	}

	wg.Wait()
	close(errCh)
	if len(errCh) > 0 {
		err := <-errCh
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File successfully uploaded and sent via gRPC"))
}

func (h *FileHandler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "Filename not specified", http.StatusBadRequest)
		return
	}

	fileMetadata, err := h.dbManager.GetFileMetadata(filename)
	if err != nil {
		http.Error(w, "Error getting file metadata: "+err.Error(), http.StatusInternalServerError)
		return
	}

	chunkMetadataList, err := h.dbManager.GetChunkMetadata(fileMetadata.ID)
	if err != nil {
		http.Error(w, "Error getting chunk metadata: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(chunkMetadataList) == 0 {
		http.Error(w, "Chunks for this file not found", http.StatusNotFound)
		return
	}

	grpcClients := h.grpcClientManager.GetClientsByName()
	if len(grpcClients) == 0 {
		http.Error(w, "No available gRPC connections", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")

	totalChunks := fileMetadata.TotalChunks

	type chunkResult struct {
		chunkNumber int32
		data        []byte
		err         error
	}

	results := make([]*chunkResult, totalChunks)

	var wg sync.WaitGroup
	for i := int32(0); i < totalChunks; i++ {
		chunkNumber := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			var metadata storage.ChunkMetadata
			found := false
			for _, meta := range chunkMetadataList {
				if meta.ChunkNumber == chunkNumber {
					metadata = meta
					found = true
					break
				}
			}
			if !found {
				results[chunkNumber] = &chunkResult{
					chunkNumber: chunkNumber,
					err:         fmt.Errorf("Metadata for chunk not found: %d", chunkNumber),
				}
				return
			}

			client := grpcClients[metadata.ServiceName]
			if client == nil {
				results[chunkNumber] = &chunkResult{
					chunkNumber: chunkNumber,
					err:         fmt.Errorf("gRPC client not found for service: %s", metadata.ServiceName),
				}
				return
			}

			chunkData, err := h.grpcClientManager.GetChunk(client, filename, metadata.ChunkNumber, metadata.ChunkHash)
			if err != nil {
				results[chunkNumber] = &chunkResult{
					chunkNumber: chunkNumber,
					err:         fmt.Errorf("Error getting chunk: %v", err),
				}
				return
			}

			chunkHash := calculateChunkHash(chunkData)
			if chunkHash != metadata.ChunkHash {
				results[chunkNumber] = &chunkResult{
					chunkNumber: chunkNumber,
					err:         fmt.Errorf("Chunk hash mismatch for chunk %d", chunkNumber),
				}
				return
			}

			results[chunkNumber] = &chunkResult{
				chunkNumber: chunkNumber,
				data:        chunkData,
				err:         nil,
			}
		}()
	}

	wg.Wait()

	for _, res := range results {
		if res.err != nil {
			http.Error(w, res.err.Error(), http.StatusInternalServerError)
			return
		}
	}

	for _, res := range results {
		_, err := w.Write(res.data)
		if err != nil {
			http.Error(w, "Error sending chunk: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Printf("File '%s' successfully downloaded\n", filename)
}
