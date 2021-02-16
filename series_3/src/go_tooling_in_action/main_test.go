package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// tdd - table driven test
	cases := []struct {
		in, out string
	}{
		{"tan@golang.org", "gopher tan"},
		{"something", "Hello, world"},
	}

	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)

		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}

		// mock response writer
		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200; got %d", rec.Code)
		}

		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("Unexpected body in response: %q", rec.Body.String())
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8000/tan@golang.org",
			nil,
		)

		if err != nil {
			b.Fatalf("Could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Expected status 200; got %d", rec.Code)
		}

		if !strings.Contains(rec.Body.String(), "gopher tan") {
			b.Errorf("Unexpected body in reqponse: %q", rec.Body.String())
		}
	}
}
