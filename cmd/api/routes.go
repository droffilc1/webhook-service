package main

import (
	"net/http"

	"github.com/droffilc1/webhook-service/internal/delivery"
	"github.com/droffilc1/webhook-service/internal/handler"
	"github.com/droffilc1/webhook-service/internal/store"
)

func (app *application) routes(st store.Store, deliveryService *delivery.DeliveryService) http.Handler {

	h := handler.New(st, deliveryService)
	mux := http.NewServeMux()

	// Unprotected routes.
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("POST /api-keys", h.CreateAPIKey)

	// Protected routes.
	protected := func(pattern string, handler http.HandlerFunc) {
		mux.Handle(pattern, apiKeyMiddleware(st, handler))
	}

	protected("POST /events", h.CreateEvent)
	protected("GET /events", h.GetEvents)
	protected("GET /events/{id}", h.GetEvent)
	protected("POST /endpoints", h.CreateEndpoint)
	protected("GET /endpoints", h.GetEndpoints)
	protected("GET /endpoints/{id}", h.GetEndpoint)
	protected("PUT /endpoints/{id}", h.UpdateEndpoint)
	protected("DELETE /endpoints/{id}", h.DeleteEndpoint)

	return securityHeaders(mux)
}
