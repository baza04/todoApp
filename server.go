package todoapp

import (
	"context"
	"net/http"
	"time"
)

// Server struct contains go http.Server ptr
type Server struct {
	httpServer *http.Server
}

// Run recieve port, handler then will start the configured server
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20, // 1MB
		ReadTimeout:       7 * time.Second,
		ReadHeaderTimeout: 7 * time.Second,
		WriteTimeout:      7 * time.Second,
		IdleTimeout:       7 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown will shutdown started server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
