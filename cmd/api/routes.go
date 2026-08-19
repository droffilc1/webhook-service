package main

import (
	"net/http"

	"github.com/droffilc1/webhook-service/internal/delivery"
	"github.com/droffilc1/webhook-service/internal/handler"
	"github.com/droffilc1/webhook-service/internal/store"
)

func (app *application) routes(st store.Store, deliveryService *delivery.DeliveryService) *http.ServeMux {

	h := handler.New(st, deliveryService)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("POST /events", h.CreateEvent)
	mux.HandleFunc("GET /events", h.GetEvents)
	mux.HandleFunc("GET /events/{id}", h.GetEvent)
	mux.HandleFunc("POST /endpoints", h.CreateEndpoint)
	mux.HandleFunc("GET /endpoints", h.GetEndpoints)
	mux.HandleFunc("GET /endpoints/{id}", h.GetEndpoint)
	mux.HandleFunc("PUT /endpoints/{id}", h.UpdateEndpoint)
	mux.HandleFunc("DELETE /endpoints/{id}", h.DeleteEndpoint)

	return mux
}
