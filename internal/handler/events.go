package handler

import (
	"encoding/json"
	"net/http"

	"github.com/droffilc1/webhook-service/internal/model"
	"github.com/droffilc1/webhook-service/internal/store"
)

type Handler struct {
	store store.Store
}

func New(store store.Store) *Handler {
	return &Handler{
		store: store,
	}
}

// CreateEvent creats a new event
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent model.Event

	err := h.store.CreateEvent(&newEvent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

// GetEvents gets the list of events
func (h *Handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	events, err := h.store.ListEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}

// GetEvent gets a specific event by its id
func (h *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	event, err := h.store.GetEvent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(event)
}
