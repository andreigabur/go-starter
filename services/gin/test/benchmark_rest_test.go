package test

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestGoGinRESTUsersEndpoint(t *testing.T) {
	const iterations = 20000

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
