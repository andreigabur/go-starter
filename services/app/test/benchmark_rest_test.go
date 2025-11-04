package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-starter-app/internal/handlers"
)

// TestRESTUsersEndpoint benchmarks the REST endpoint performance
func TestRESTUsersEndpoint(t *testing.T) {
	const iterations = 10000
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
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
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

// TestRESTUsersEndpointWithBodyRead benchmarks including reading the response body
func TestRESTUsersEndpointWithBodyRead(t *testing.T) {
	const iterations = 10000
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
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Read the body to simulate full request/response cycle
		body := rr.Body.Bytes()
		if len(body) == 0 {
			t.Error("expected non-empty response body")
		}
	}
	elapsed := time.Since(start)

	avgTime := elapsed / iterations
	reqPerSec := float64(time.Second) / float64(avgTime)

	t.Logf("REST /users endpoint (with body read):")
	t.Logf("  Iterations: %d", iterations)
	t.Logf("  Total time: %v", elapsed)
	t.Logf("  Time per request: %v", avgTime)
	t.Logf("  Requests per second: %.0f", reqPerSec)
}
