package test

import (
	"context"
	"testing"
	"time"

	"go-starter-app/internal/grpc"
	"go-starter-app/pkg/pb"
)

// BenchmarkGRPCUsers benchmarks the gRPC endpoint performance
func BenchmarkGRPCUsers(b *testing.B) {
	service := grpc.NewUsersService()
	ctx := context.Background()
	req := &pb.ListUsersRequest{}

	b.ResetTimer()
	b.ReportAllocs()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		resp, err := service.ListUsers(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
		if resp == nil {
			b.Fatal("expected non-nil response")
		}
		if len(resp.Users) == 0 {
			b.Error("expected non-empty users list")
		}
	}
}

// BenchmarkGRPCUsersWithConversion benchmarks including full conversion overhead
func BenchmarkGRPCUsersWithConversion(b *testing.B) {
	service := grpc.NewUsersService()
	ctx := context.Background()
	req := &pb.ListUsersRequest{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		resp, err := service.ListUsers(ctx, req)
		if err != nil {
			b.Fatal(err)
		}

		// Access the converted data to simulate full processing
		for _, user := range resp.Users {
			_ = user.Id
			_ = user.Name
			_ = user.Email
		}
	}
}

// TestGRPCUsersSpeed runs a performance test and reports timing
func TestGRPCUsersSpeed(t *testing.T) {
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
	const iterations = 10000
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_, err := service.ListUsers(ctx, req)
		if err != nil {
			t.Fatal(err)
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
