package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
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
