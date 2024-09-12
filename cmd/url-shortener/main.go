package main

import (
	"RESTFullGolang/internal/config"
	"RESTFullGolang/internal/logger"
	"log/slog"
)

// config, logger, storage, router, server
func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info("start", slog.String("environment", cfg.Env))
	log.Debug("debugging")
}
