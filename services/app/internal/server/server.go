package server

import (
	"go-starter-app/internal/handlers"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.mux.HandleFunc("/health", handlers.HealthCheck)
	s.mux.HandleFunc("/users", handlers.HandleUsers)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
