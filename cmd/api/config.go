package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	addr string
	dsn  string
}

func loadConfig() config {
	godotenv.Load()

	cfg := config{
		addr: os.Getenv("PORT"),
		dsn:  os.Getenv("DATABASE_URL"),
	}
	if cfg.addr != "" {
		cfg.addr = ":" + cfg.addr
	}
	if cfg.addr == "" {
		cfg.addr = ":4000"
	}
	if cfg.dsn == "" {
		slog.Error("DATABASE_URL is required")
	}

	return cfg
}
