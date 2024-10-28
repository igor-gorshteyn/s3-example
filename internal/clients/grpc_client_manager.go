package clients

import (
	"context"
	"errors"
	"sync"
	"time"

	filetransfer "s3-example/api/gen/go"

	"google.golang.org/grpc"
)

type GrpcClientManager struct {
	mu      sync.RWMutex
	clients map[string]filetransfer.FileTransferServiceClient
}

func NewGrpcClientManager() *GrpcClientManager {
	return &GrpcClientManager{
		clients: make(map[string]filetransfer.FileTransferServiceClient),
	}
}

func (m *GrpcClientManager) RegisterClient(serviceName, address string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.clients[serviceName]; exists {
		return errors.New("gRPC client already registered")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return err
	}

	client := filetransfer.NewFileTransferServiceClient(conn)
	m.clients[serviceName] = client

	return nil
}

func (m *GrpcClientManager) GetClients() []filetransfer.FileTransferServiceClient {
	m.mu.RLock()
	defer m.mu.RUnlock()

	clients := make([]filetransfer.FileTransferServiceClient, 0, len(m.clients))
	for _, client := range m.clients {
		clients = append(clients, client)
	}

	return clients
}

func (m *GrpcClientManager) GetClientNames() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	names := make([]string, 0, len(m.clients))
	for name := range m.clients {
		names = append(names, name)
	}

	return names
}

func (m *GrpcClientManager) GetClientsByName() map[string]filetransfer.FileTransferServiceClient {
	m.mu.RLock()
	defer m.mu.RUnlock()

	clientsCopy := make(map[string]filetransfer.FileTransferServiceClient)
	for name, client := range m.clients {
		clientsCopy[name] = client
	}

	return clientsCopy
}

func (m *GrpcClientManager) GetClientByName(name string) filetransfer.FileTransferServiceClient {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.clients[name]
}

func (m *GrpcClientManager) SendChunks(client filetransfer.FileTransferServiceClient, chunks []*filetransfer.FileChunk) error {
	ctx := context.Background()
	stream, err := client.TransferFile(ctx)
	if err != nil {
		return err
	}

	for _, chunk := range chunks {
		if err := stream.Send(chunk); err != nil {
			return err
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}

	return nil
}

func (m *GrpcClientManager) GetChunk(client filetransfer.FileTransferServiceClient, filename string, chunkNumber int32, chunkHash string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := &filetransfer.ChunkRequest{
		Filename:    filename,
		ChunkNumber: chunkNumber,
		ChunkHash:   chunkHash,
	}

	response, err := client.GetChunk(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.Chunk, nil
}
