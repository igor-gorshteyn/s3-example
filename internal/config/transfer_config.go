package config

type TransferServiceConfig struct {
	ServerPort       string
	GRPCPort         string
	RedisAddr        string
	SessionTTL       int
	MaxUploadSize    int64
	ChunkSize        int
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

func LoadTransferConfig() (*TransferServiceConfig, error) {
	return &TransferServiceConfig{
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		GRPCPort:         getEnv("GRPC_PORT", "5001"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		SessionTTL:       getEnvAsInt("SESSION_TTL", 3600),
		MaxUploadSize:    getEnvAsInt64("MAX_UPLOAD_SIZE_GB", 2) * 1024 * 1024 * 1024,
		ChunkSize:        int(getEnvAsInt64("CHUNK_SIZE_BYTES", 1048576)),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "user"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),
		PostgresDBName:   getEnv("POSTGRES_DB", "dbname"),
	}, nil
}
