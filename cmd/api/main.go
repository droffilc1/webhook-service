package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/droffilc1/webhook-service/internal/delivery"
	"github.com/droffilc1/webhook-service/internal/handler"
	"github.com/droffilc1/webhook-service/internal/store"
)

func main() {
	s := store.NewStore()

	deliveryService := delivery.New(s, &http.Client{
		Timeout: 5 * time.Second,
	})

	h := handler.New(s, deliveryService)

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	logger.Info("starting server", "addr", *addr)
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
