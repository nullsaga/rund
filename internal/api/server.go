package api

import (
	"context"
	"github.com/nullsaga/rund/internal/conf"
	"log/slog"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

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

func (s *Server) RegisterHandlers(rootConf *conf.RootConf) {
	s.router.HandleFunc("POST /v1/webhook/{hook}", s.makeHandler(NewWebhookHandler().Handle))
}

func (s *Server) Start() error {
	slog.Info("starting api server", "addr", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	slog.Info("attempting to shutdown down api server")
	return s.server.Shutdown(ctx)
}

func (s *Server) makeHandler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			slog.Error(err.Error())
		}
	}
}
