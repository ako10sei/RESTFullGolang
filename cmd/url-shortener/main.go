package main

import (
	"RESTFullGolang/internal/config"
	"RESTFullGolang/internal/http-server/middleware"
	"RESTFullGolang/internal/lib/logger/sl"
	"RESTFullGolang/internal/logger"
	"RESTFullGolang/internal/storage/sqlite"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
)

// config, logger, storage, router, server
func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)
	log.Info("start", slog.String("environment", cfg.Env))
	log.Debug("debugging")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("error creating storage", sl.Err(err))
		os.Exit(1)
	}
	defer func(storage *sqlite.Storage) {
		err = storage.Close()
		if err != nil {
			log.Error("error closing storage", sl.Err(err))
		}
	}(storage)

	router := chi.NewRouter()
	middleware.AppendMiddleware(router, log)

}
