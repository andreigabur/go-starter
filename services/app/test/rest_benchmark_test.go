package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-starter-app/internal/handlers"
)

// BenchmarkRESTUsers benchmarks the REST endpoint performance
func BenchmarkRESTUsers(b *testing.B) {
	// Create a test request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create a fresh ResponseRecorder for each iteration
		rr := httptest.NewRecorder()

		// Execute the handler
		handlers.HandleUsers(rr, req)

		// Check status code
		if status := rr.Code; status != http.StatusOK {
			b.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	}
}

// BenchmarkRESTUsersWithBodyRead benchmarks including reading the response body
func BenchmarkRESTUsersWithBodyRead(b *testing.B) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		rr := httptest.NewRecorder()
		handlers.HandleUsers(rr, req)

		// Read the body to simulate full request/response cycle
		body := rr.Body.Bytes()
		if len(body) == 0 {
			b.Error("expected non-empty response body")
		}
	}
}

// TestRESTUsersSpeed runs a performance test and reports timing
func TestRESTUsersSpeed(t *testing.T) {
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
	const iterations = 10000
	start := time.Now()
	for i := 0; i < iterations; i++ {
		rr := httptest.NewRecorder()
		handlers.HandleUsers(rr, req)
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
