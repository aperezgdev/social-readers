package main

import (
	"log/slog"
	"net/http"

	"github.com/aperezgdev/social-readers-api/internal/infrastructure/config"
	server "github.com/aperezgdev/social-readers-api/internal/infrastructure/http"
)

func main() {
	slog := slog.Default()
	config := config.NewConfig(slog)
	httpServer := server.NewHttpServer(slog, config)

	httpServer.Handler().HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	err := httpServer.Start()
	if err != nil {
		slog.Error("Error starting server")
		panic(err)
	}
}
