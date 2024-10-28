GO ?= go
PROTOC ?= protoc

API_PROTO_DIR := api/proto
GEN_GO_DIR := api/gen/go
PROTO_FILES := $(API_PROTO_DIR)/*.proto

.PHONY: proto
proto:
	@echo "Generating gRPC code..."
	@mkdir -p $(GEN_GO_DIR)
	$(PROTOC) -I $(API_PROTO_DIR) \
		--go_out=$(GEN_GO_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_GO_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
	@echo "Generation completed."

.PHONY: build-transfer
build-transfer: proto
	@echo "Building transferService..."
	$(GO) build -o bin/transferService ./cmd/transferService
	@echo "transferService build completed."

.PHONY: build-storage
build-storage: proto
	@echo "Building storageService..."
	$(GO) build -o bin/storageService ./cmd/storageService
	@echo "storageService build completed."

.PHONY: build
build: build-transfer build-storage

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(GEN_GO_DIR) bin
	@echo "Clean up completed."


.PHONY: migrate
migrate:
	@goose -dir ./migrations postgres "host=localhost user=yourusername password=yourpassword dbname=yourdbname sslmode=disable" up
