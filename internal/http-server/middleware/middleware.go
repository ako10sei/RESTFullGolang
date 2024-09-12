package middleware

import (
	mwLogger "RESTFullGolang/internal/http-server/middleware/logger"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AppendMiddleware(router *chi.Mux, log *slog.Logger) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}
