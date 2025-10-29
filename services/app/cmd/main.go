package main

import (
	"go-starter-app/internal/server"
	"log"
)

func main() {
	srv := server.NewServer()

	log.Println("Starting server on :8080")
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
