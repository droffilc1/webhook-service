package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/droffilc1/webhook-service/internal/delivery"
	"github.com/droffilc1/webhook-service/internal/store"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	logger *slog.Logger
}

func main() {

	cfg := loadConfig()
	addr := flag.String("addr", cfg.addr, "HTTP network address")
	dsn := flag.String(
		"dsn",
		cfg.dsn,
		"Postgres data source name",
	)
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer func() {
		if cerr := db.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	var st store.Store
	if *dsn != "" {
		st = store.NewPostgresStore(db)
	} else {
		st = store.NewStore()
	}

	deliveryService := delivery.New(st, &http.Client{
		Timeout: 5 * time.Second,
	})

	app := &application{
		logger: logger,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(st, deliveryService),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("starting server", "addr", srv.Addr)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		defer func() {
			if cerr := db.Close(); cerr != nil && err == nil {
				err = cerr
			}
		}()
		return nil, err
	}

	return db, nil
}
