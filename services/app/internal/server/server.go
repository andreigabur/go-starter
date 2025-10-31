package server

import (
	"log"
	"net"
	"net/http"

	"go-starter-app/internal/grpc"
	"go-starter-app/internal/handlers"
	"go-starter-app/pkg/pb"

	grpcpkg "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (s *Server) startGRPC(addr string) {
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen on %s: %v", addr, err)
		}
		gs := grpcpkg.NewServer()
		pb.RegisterUsersServiceServer(gs, grpc.NewUsersService())
		reflection.Register(gs)
		log.Printf("gRPC server listening on %s", addr)
		if err := gs.Serve(lis); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()
}

func (s *Server) Start(addr string) error {
	// Start gRPC on :9090 in background
	s.startGRPC(":9090")
	return http.ListenAndServe(addr, s.mux)
}
