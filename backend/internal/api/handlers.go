package api

import (
	"encoding/json"
	"net/http"

	"github.com/bytepoet/hasbowleh/internal/database"
	"github.com/bytepoet/hasbowleh/internal/v2ray"
	"github.com/bytepoet/hasbowleh/internal/wireguard"
	"github.com/gorilla/mux"
)

type Handler struct {
	DB *database.DB
	WG *wireguard.WireGuard
	V2 *v2ray.V2Ray
}

func (h *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := h.DB.GetClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(clients)
}

func (h *Handler) AddClient(w http.ResponseWriter, r *http.Request) {
	var client database.Client
	json.NewDecoder(r.Body).Decode(&client)

	// Generate WireGuard and v2ray configurations
	wgConfig, err := h.WG.GenerateClientConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v2Config, err := h.V2.GenerateClientConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save client to database
	client.WireGuardConfig = wgConfig
	client.V2RayConfig = v2Config
	err = h.DB.AddClient(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func (h *Handler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID := vars["id"]

	err := h.DB.DeleteClient(clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Remove client from WireGuard and v2ray
	h.WG.RemoveClient(clientID)
	h.V2.RemoveClient(clientID)

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DownloadConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID := vars["id"]

	client, err := h.DB.GetClient(clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Combine WireGuard and v2ray configs
	config := client.WireGuardConfig + "\n" + client.V2RayConfig

	w.Header().Set("Content-Disposition", "attachment; filename=config.conf")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(config))
}
