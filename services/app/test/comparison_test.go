package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-starter-app/internal/grpc"
	"go-starter-app/internal/handlers"
	"go-starter-app/pkg/pb"
)

// TestRESTvsGRPCComparison compares REST and gRPC endpoint performance side-by-side
func TestRESTvsGRPCComparison(t *testing.T) {
	const iterations = 10000

	// Test REST endpoint
	t.Run("REST", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Warm up
		for i := 0; i < 10; i++ {
			rr := httptest.NewRecorder()
			handlers.HandleUsers(rr, req)
		}

		// Measure
		start := time.Now()
		for i := 0; i < iterations; i++ {
			rr := httptest.NewRecorder()
			handlers.HandleUsers(rr, req)
		}
		restElapsed := time.Since(start)

		restAvgTime := restElapsed / iterations
		restReqPerSec := float64(time.Second) / float64(restAvgTime)

		t.Logf("REST /users endpoint:")
		t.Logf("  Iterations: %d", iterations)
		t.Logf("  Total time: %v", restElapsed)
		t.Logf("  Time per request: %v", restAvgTime)
		t.Logf("  Requests per second: %.0f", restReqPerSec)
	})

	// Test gRPC endpoint
	t.Run("gRPC", func(t *testing.T) {
		service := grpc.NewUsersService()
		ctx := context.Background()
		req := &pb.ListUsersRequest{}

		// Warm up
		for i := 0; i < 10; i++ {
			_, err := service.ListUsers(ctx, req)
			if err != nil {
				t.Fatal(err)
			}
		}

		// Measure
		start := time.Now()
		for i := 0; i < iterations; i++ {
			_, err := service.ListUsers(ctx, req)
			if err != nil {
				t.Fatal(err)
			}
		}
		grpcElapsed := time.Since(start)

		grpcAvgTime := grpcElapsed / iterations
		grpcReqPerSec := float64(time.Second) / float64(grpcAvgTime)

		t.Logf("gRPC ListUsers endpoint:")
		t.Logf("  Iterations: %d", iterations)
		t.Logf("  Total time: %v", grpcElapsed)
		t.Logf("  Time per request: %v", grpcAvgTime)
		t.Logf("  Requests per second: %.0f", grpcReqPerSec)
	})
}

