package todoapp

import (
	"context"
	"net/http"
	"time"
)

const timeout = 7

// Server struct contains go http.Server ptr.
type Server struct {
	httpServer *http.Server
}

// Run receive port, handler then will start the configured server.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20, // 1MB
		ReadTimeout:       timeout * time.Second,
		ReadHeaderTimeout: timeout * time.Second,
		WriteTimeout:      timeout * time.Second,
		IdleTimeout:       timeout * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown will shut down started server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
