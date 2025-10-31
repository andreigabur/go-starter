package controllers

import "go-starter-app/internal/models"

// ListUsers returns a list of hardcoded users.
// This is the shared business logic used by both REST and gRPC handlers.
func ListUsers() []models.User {
	return []models.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}
}
