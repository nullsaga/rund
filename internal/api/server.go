package api

import (
	"context"
	"log/slog"
	"net/http"
)

type Server struct {
	router *http.ServeMux
	server *http.Server
}

func NewServer(addr string) *Server {
	router := http.NewServeMux()
	return &Server{
		router: router,
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) RegisterHandlers() {
	// TODO
}

func (s *Server) Start() error {
	slog.Info("starting api server", "addr", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	slog.Info("attempting to shutdown down api server")
	return s.server.Shutdown(ctx)
}
