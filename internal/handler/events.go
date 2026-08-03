package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/droffilc1/webhook-service/internal/delivery"
	"github.com/droffilc1/webhook-service/internal/model"
	"github.com/droffilc1/webhook-service/internal/store"
	"github.com/google/uuid"
)

type Handler struct {
	store    store.Store
	delivery *delivery.DeliveryService
}

func New(store store.Store, delivery *delivery.DeliveryService) *Handler {
	return &Handler{
		store:    store,
		delivery: delivery,
	}
}

// CreateEvent creats a new event
func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent model.Event

	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newEvent.ID = uuid.NewString()
	newEvent.CreatedAt = time.Now()

	err := h.delivery.DeliverEvent(&newEvent)
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
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(event)
}
