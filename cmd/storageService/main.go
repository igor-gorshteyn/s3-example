package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"s3-example/internal/config"
	"s3-example/internal/server"
)

func main() {
	cfg, err := config.LoadStorageConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	storageDir := cfg.StorageDir

	err = os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating storage directory: %v", err)
	}

	go func() {
		err := server.StartStorageGRPCServer(cfg.GRPCPort, storageDir, cfg.ServiceName)
		if err != nil {
			log.Fatalf("Error starting gRPC server: %v", err)
		}
	}()

	err = registerWithTransferService(cfg)
	if err != nil {
		log.Fatalf("Error registering with Transfer Service: %v", err)
	}

	select {}
}

func registerWithTransferService(cfg *config.StorageServiceConfig) error {
	grpcAddress := fmt.Sprintf("%s:%s", cfg.ServiceName, cfg.GRPCPort)

	reqBody, err := json.Marshal(map[string]string{
		"service_name": cfg.ServiceName,
		"grpc_address": grpcAddress,
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(cfg.TransferServiceURL+"/register", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register, status code: %d", resp.StatusCode)
	}

	log.Printf("Successfully registered with Transfer Service as %s at %s", cfg.ServiceName, grpcAddress)
	return nil
}
