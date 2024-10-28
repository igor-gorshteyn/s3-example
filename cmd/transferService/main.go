package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"s3-example/internal/clients"
	"s3-example/internal/config"
	"s3-example/internal/handlers"
	"s3-example/internal/storage"

	"github.com/pressly/goose/v3"
)

func main() {
	cfg, err := config.LoadTransferConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	var dbManager *storage.Manager
	for {
		dbManager, err = storage.NewDBManager(cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDBName)
		if err != nil {
			log.Printf("Error connecting to database: %v", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	defer dbManager.Close()

	if err := goose.Up(dbManager.DB, "migrations"); err != nil {
		log.Fatalf("Error applying migrations: %v", err)
	}

	grpcClientManager := clients.NewGrpcClientManager()

	fileHandler := handlers.NewFileHandler(cfg, grpcClientManager, dbManager)
	registrationHandler := handlers.NewRegistrationHandler(grpcClientManager)

	http.HandleFunc("/upload", fileHandler.UploadHandler)
	http.HandleFunc("/download", fileHandler.DownloadHandler)

	http.HandleFunc("/register", registrationHandler.RegisterHandler)
	http.HandleFunc("/clients", registrationHandler.GetClientsHandler)

	fmt.Printf("HTTP server started on port %s\n", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
