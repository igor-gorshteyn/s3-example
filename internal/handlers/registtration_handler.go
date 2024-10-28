package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"s3-example/internal/clients"
)

type RegistrationHandler struct {
	grpcClientManager *clients.GrpcClientManager
}

func NewRegistrationHandler(grpcClientManager *clients.GrpcClientManager) *RegistrationHandler {
	return &RegistrationHandler{
		grpcClientManager: grpcClientManager,
	}
}

func (h *RegistrationHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ServiceName string `json:"service_name"`
		GRPCAddress string `json:"grpc_address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.grpcClientManager.RegisterClient(req.ServiceName, req.GRPCAddress); err != nil {
		http.Error(w, "Failed to register client: "+err.Error(), http.StatusBadRequest)
		return
	}

	clientAddresses := h.grpcClientManager.GetClientNames()
	log.Printf("Registered new client: %s. Total clients: %v", req.ServiceName, clientAddresses)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Client registered successfully"))
}

func (h *RegistrationHandler) GetClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients := h.grpcClientManager.GetClientNames()
	response, err := json.Marshal(clients)
	if err != nil {
		http.Error(w, "Error forming response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
