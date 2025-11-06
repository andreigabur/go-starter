package test

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"go-starter-app/internal/handlers"
)

// startTestServer starts the actual HTTP server on a dynamic port using the
// application's mux. It returns the base URL (e.g. http://127.0.0.1:12345)
// and a stop function to shut the server down.
func startTestServer(t *testing.T) (string, func()) {
	t.Helper()
	// Build a mux using the existing handlers to avoid depending on the
	// server.Server internal implementation (mux is unexported).
	mux := http.NewServeMux()
	mux.HandleFunc("/users", handlers.HandleUsers)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen on a port: %v", err)
	}

	srv := &http.Server{Handler: mux}

	go func() {
		// Serve will return when ln is closed or on error
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			t.Logf("test server serve error: %v", err)
		}
	}()

	baseURL := fmt.Sprintf("http://%s", ln.Addr().String())

	stop := func() {
		// give 1s to shutdown
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}
	return baseURL, stop
}

// TestRESTUsersEndpoint benchmarks the REST endpoint performance by hitting
// the real HTTP server and reading the response body
func TestGoStandardRESTUsersEndpoint(t *testing.T) {
	const iterations = 20000

	// baseURL, stop := startTestServer(t)
	// defer stop()
	baseURL := "http://localhost:8080"

	url := baseURL + "/users"

	// Warm up
	for i := 0; i < 10; i++ {
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("warmup request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("warmup returned status %d", resp.StatusCode)
		}
		_, _ = io.ReadAll(resp.Body)
		_ = resp.Body.Close()
	}

	// Measure
	start := time.Now()
	for i := 0; i < iterations; i++ {
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
		}

		body, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}
		if len(body) == 0 {
			t.Fatalf("expected non-empty response body")
		}
	}
	elapsed := time.Since(start)

	avgTime := elapsed / iterations
	reqPerSec := float64(time.Second) / float64(avgTime)

	t.Logf("REST /users endpoint:")
	t.Logf("  Iterations: %d", iterations)
	t.Logf("  Total time: %v", elapsed)
	t.Logf("  Time per request: %v", avgTime)
	t.Logf("  Requests per second: %.0f", reqPerSec)
}
