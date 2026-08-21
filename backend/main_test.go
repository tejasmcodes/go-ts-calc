package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMathHandler(t *testing.T) {
	// create a mock http request
	req := httptest.NewRequest(http.MethodGet, "/calculate?a=1&b=2&operation=add", nil)

	// create a response recorder to capture the response
	rec := httptest.NewRecorder()

	// call the math handler directly
	mathHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected JSON content type")
	}

	var response MathResponse

	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decoded response: %v", err)
	}

	if response.A != 1 {
		t.Errorf("expected A=1, got %v", response.A)
	}

	if response.B != 2 {
		t.Errorf("expected B=2, got %v", response.B)
	}

	if response.Operation != "add" {
		t.Errorf("Expected operation = add, got %v", response.Operation)
	}

	if response.Result != 3 {
		t.Errorf("Expected result=3, got %v", response.Result)
	}
}
