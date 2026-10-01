package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/droffilc1/webhook-service/internal/auth"
	"github.com/droffilc1/webhook-service/internal/model"
	"github.com/google/uuid"
)

type createAPIKeyRequest struct {
	Name string `json:"name"`
}

type createAPIKeyResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// CreateAPIKey creates a new API key.
func (h *Handler) CreateAPIKey(w http.ResponseWriter, r *http.Request) {
	var req createAPIKeyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, hash := auth.Generate()

	now := time.Now()
	expires := now.Add(90 * 24 * time.Hour)

	apiKey := &model.APIKey{
		ID:        uuid.NewString(),
		KeyHash:   hash,
		Name:      req.Name,
		CreatedAt: now,
		ExpiresAt: &expires,
	}

	if err := h.store.CreateAPIKey(apiKey); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := createAPIKeyResponse{
		ID:        apiKey.ID,
		Name:      apiKey.Name,
		Key:       key,
		CreatedAt: apiKey.CreatedAt,
		ExpiresAt: *apiKey.ExpiresAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
