package todoapp

import (
	"context"
	"net/http"
	"time"

	"github.com/baza04/todoApp/pkg/handler"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler handler.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    7 * time.Second,
		WriteTimeout:   7 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
