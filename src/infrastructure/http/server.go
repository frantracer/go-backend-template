package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	appHandlers "github.com/frantacer/go-backend-template/src/application/handlers"
)

type Server struct {
	server        *http.Server
	serverContext context.Context
}

func NewServer(ctx context.Context, handler http.Handler) Server {
	server := &http.Server{
		Handler: handler,
	}

	server.BaseContext = func(ln net.Listener) context.Context {
		return ctx
	}

	return Server{server: server, serverContext: ctx}
}

func (s *Server) ListenAndServe(port int, readyCh chan struct{}) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	close(readyCh)

	go func() {
		<-s.serverContext.Done()
		_ = s.server.Shutdown(s.serverContext)
	}()

	err = s.server.Serve(listener)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

type ApplicationHandlers struct {
	FindTasksHandler  appHandlers.FindTasksHandler
	InsertTaskHandler appHandlers.InsertTaskHandler
}

// NewHandler defines the middlewares used and the endpoints exposed
func NewHandler(appHandlers ApplicationHandlers) http.Handler {
	const requestTimeoutInSeconds = 60

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Timeout(requestTimeoutInSeconds * time.Second))

	r.Get("/health", HealthCheckHTTPFunc())

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", FindTasksHTTPFunc(appHandlers.FindTasksHandler))
		r.Post("/", InsertTaskHTTPFunc(appHandlers.InsertTaskHandler))
	})

	return r
}
