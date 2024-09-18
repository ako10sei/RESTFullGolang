package main

import (
	"RESTFullGolang/internal/config"
	"RESTFullGolang/internal/http-server/handlers/redirect"
	"RESTFullGolang/internal/http-server/handlers/remove"
	"RESTFullGolang/internal/http-server/handlers/url/save"
	"RESTFullGolang/internal/http-server/middleware"
	"RESTFullGolang/internal/lib/logger/sl"
	"RESTFullGolang/internal/logger"
	"RESTFullGolang/internal/storage/sqlite"
	"log/slog"
	"net/http"
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

	router.Post("/url", save.New(log, storage))
	router.Get("/{alias}", redirect.New(log, storage))
	router.Delete("/{alias}", remove.New(log, storage))

	log.Info("start server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")

}
