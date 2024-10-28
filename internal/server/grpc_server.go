package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	filetransfer "s3-example/api/gen/go"

	"google.golang.org/grpc"
)

const filesPath = "files"

type FileTransferServer struct {
	filetransfer.UnimplementedFileTransferServiceServer
	StorageDir  string
	ServiceName string
}

func NewFileTransferServer(storageDir, serviceName string) *FileTransferServer {
	return &FileTransferServer{
		StorageDir:  storageDir,
		ServiceName: serviceName,
	}
}

func (s *FileTransferServer) TransferFile(stream filetransfer.FileTransferService_TransferFileServer) error {
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&filetransfer.TransferResponse{
				Status: "File successfully received",
			})
		}
		if err != nil {
			return err
		}

		err = s.saveChunk(chunk)
		if err != nil {
			return err
		}
	}
}

func (s *FileTransferServer) saveChunk(chunk *filetransfer.FileChunk) error {
	fileDir := filepath.Join(s.StorageDir, filesPath, s.ServiceName)
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		return err
	}

	chunkFilename := chunk.ChunkHash

	chunkPath := filepath.Join(fileDir, chunkFilename)

	err = os.WriteFile(chunkPath, chunk.Chunk, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (s *FileTransferServer) GetChunk(ctx context.Context, req *filetransfer.ChunkRequest) (*filetransfer.ChunkResponse, error) {
	fileDir := filepath.Join(s.StorageDir, filesPath, s.ServiceName)

	chunkHash := req.ChunkHash
	chunkPath := filepath.Join(fileDir, chunkHash)

	if _, err := os.Stat(chunkPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Chunk not found")
	}

	chunkData, err := os.ReadFile(chunkPath)
	if err != nil {
		return nil, err
	}

	return &filetransfer.ChunkResponse{
		Chunk: chunkData,
	}, nil
}

func StartStorageGRPCServer(port string, storageDir string, serviceName string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	ftServer := NewFileTransferServer(storageDir, serviceName)
	filetransfer.RegisterFileTransferServiceServer(server, ftServer)

	log.Printf("StorageService '%s' gRPC server started on port %s", serviceName, port)
	return server.Serve(listener)
}
