// Package handler implements handler functions.
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/droffilc1/webhook-service/internal/model"
	"github.com/google/uuid"
)

// CreateEndpoint creates a new endpoint
func (h *Handler) CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	var newEndpoint model.Endpoint

	if err := json.NewDecoder(r.Body).Decode(&newEndpoint); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newEndpoint.ID = uuid.NewString()
	newEndpoint.CreatedAt = time.Now()

	err := h.store.CreateEndpoint(&newEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newEndpoint); err != nil {
		log.Fatal(err)
	}
}

// GetEndpoints gets a list of endpoints
func (h *Handler) GetEndpoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpoints, err := h.store.ListEndpoints()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(endpoints); err != nil {
		log.Fatal(err)
	}
}

// GetEndpoint gets a specific event by its id
func (h *Handler) GetEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	endpoint, err := h.store.GetEndpoint(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(endpoint); err != nil {
		log.Fatal(err)
	}
}

// UpdateEndpoint updates contents of the endpoint
func (h *Handler) UpdateEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	var endpoint model.Endpoint
	endpoint.ID = id

	err := h.store.UpdateEndpoint(&endpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(endpoint); err != nil {
		log.Fatal(err)
	}
}

// DeleteEndpoint deletes an endpoint
func (h *Handler) DeleteEndpoint(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.store.DeleteEndpoint(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
