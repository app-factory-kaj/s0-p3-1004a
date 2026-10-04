package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGreetingWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Ada", nil)
	rec := httptest.NewRecorder()

	GetGreeting(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var g Greeting
	if err := json.NewDecoder(rec.Body).Decode(&g); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if g.Message != "Hello, Ada!" {
		t.Fatalf("expected %q, got %q", "Hello, Ada!", g.Message)
	}
}

func TestGetGreetingDefaultsToWorld(t *testing.T) {
	for _, url := range []string{"/hello", "/hello?name="} {
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rec := httptest.NewRecorder()

		GetGreeting(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rec.Code)
		}
		var g Greeting
		if err := json.NewDecoder(rec.Body).Decode(&g); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if g.Message != "Hello, World!" {
			t.Fatalf("expected %q, got %q", "Hello, World!", g.Message)
		}
	}
}
