package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandlerReturnsOK(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	NewRouter().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	contentType := response.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Fatalf("expected json content type, got %s", contentType)
	}

	var body map[string]string

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("expected status ok, got %s", body["status"])
	}

	if body["service"] != "tasks" {
		t.Fatalf("expected service tasks, got %s", body["service"])
	}
}

func TestHealthHandlerRejectsPostMethod(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/health", nil)
	response := httptest.NewRecorder()

	NewRouter().ServeHTTP(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, response.Code)
	}
}