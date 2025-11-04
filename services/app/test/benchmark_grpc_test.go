package test

import (
	"context"
	"testing"
	"time"

	"go-starter-app/internal/grpc"
	"go-starter-app/pkg/pb"
)

// TestGRPCUsersEndpoint benchmarks the gRPC endpoint performance
func TestGRPCUsersEndpoint(t *testing.T) {
	const iterations = 10000
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
		resp, err := service.ListUsers(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}
		if len(resp.Users) == 0 {
			t.Error("expected non-empty users list")
		}
	}
	elapsed := time.Since(start)

	avgTime := elapsed / iterations
	reqPerSec := float64(time.Second) / float64(avgTime)

	t.Logf("gRPC ListUsers endpoint:")
	t.Logf("  Iterations: %d", iterations)
	t.Logf("  Total time: %v", elapsed)
	t.Logf("  Time per request: %v", avgTime)
	t.Logf("  Requests per second: %.0f", reqPerSec)
}

// TestGRPCUsersEndpointWithConversion benchmarks including full conversion overhead
func TestGRPCUsersEndpointWithConversion(t *testing.T) {
	const iterations = 10000
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
		resp, err := service.ListUsers(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			t.Fatal("expected non-nil response")
		}

		// Access the converted data to simulate full processing
		for _, user := range resp.Users {
			_ = user.Id
			_ = user.Name
			_ = user.Email
		}
	}
	elapsed := time.Since(start)

	avgTime := elapsed / iterations
	reqPerSec := float64(time.Second) / float64(avgTime)

	t.Logf("gRPC ListUsers endpoint (with conversion):")
	t.Logf("  Iterations: %d", iterations)
	t.Logf("  Total time: %v", elapsed)
	t.Logf("  Time per request: %v", avgTime)
	t.Logf("  Requests per second: %.0f", reqPerSec)
}
