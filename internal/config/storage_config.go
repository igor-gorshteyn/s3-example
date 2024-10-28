package config

type StorageServiceConfig struct {
	GRPCPort           string
	TransferServiceURL string
	StorageDir         string
	ServiceName        string
}

func LoadStorageConfig() (*StorageServiceConfig, error) {
	grpcPort := getEnv("GRPC_PORT", "5001")
	transferServiceURL := getEnv("TRANSFER_SERVICE_URL", "http://transfer_service:8080")
	storageDir := getEnv("STORAGE_DIR", "./storage")
	serviceName := getEnv("SERVICE_NAME", "default_service_name")

	return &StorageServiceConfig{
		GRPCPort:           grpcPort,
		TransferServiceURL: transferServiceURL,
		StorageDir:         storageDir,
		ServiceName:        serviceName,
	}, nil
}
