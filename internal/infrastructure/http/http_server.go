package server

import (
	"log/slog"
	"net/http"

	"github.com/aperezgdev/social-readers-api/internal/infrastructure/config"
)

type HttpServer struct {
	slog       *slog.Logger
	port       string
	httpServer *http.Server
	handler    *http.ServeMux
}

func NewHttpServer(slog *slog.Logger, config config.Config) *HttpServer {
	handler := http.NewServeMux()

	server := http.Server{
		Handler: handler,
		Addr:    ":" + config.ServerPort,
	}

	return &HttpServer{
		port:       config.ServerPort,
		slog:       slog,
		httpServer: &server,
		handler:    handler,
	}
}

func (hs *HttpServer) Handler() *http.ServeMux {
	return hs.handler
}

func (hs *HttpServer) Start() error {
	return hs.httpServer.ListenAndServe()
}