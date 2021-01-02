package http

import (
	"github.com/frantacer/go-backend-template/src/infrastructure/http/functions"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const requestTimeoutInSeconds = 60

// NewHandler defines the middlewares used and the endpoints exposed
func NewHandler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Timeout(requestTimeoutInSeconds * time.Second))

	// Endpoints definition
	r.Get("/health", functions.HealthCheckHandler())

	return r
}
