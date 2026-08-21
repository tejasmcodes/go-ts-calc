package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)
// def func to test the success
func TestMathHandlerSuccess(t *testing.T){
	tests := [] struct {
	name string
	a string
	b string
	operation string
	expected float64
	}{
		{
			name:	"addition",
			a:		"10",
			b:		"5",
			operation:"add",
			expected: 15,
	},
	{
			name:	"subtraction",
			a:		"10",
			b:		"5",
			operation:"sub",
			expected: 5,
	},
	{
			name:	"multiplication",
			a:		"10",
			b:		"5",
			operation:"mul",
			expected: 50,
	},
	{
			name:	"division",
			a:		"10",
			b:		"5",
			operation:"div",
			expected: 2,
	},
	}

	for _, test := range tests {
	t.Run(test.name, func(t *testing.T){
		url := fmt.Sprintf(
			"/calculate?a=%s&b=%s&operation=%s",
			test.a,
			test.b,
			test.operation,
		)
		req := httptest.NewRequest(http.MethodGet, url, nil)
		rec := httptest.NewRecorder()

		mathHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf(
				"Expected status %d, got %d", 
				http.StatusOK,
				rec.Code,

			)
		}

		var response MathResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to decode the response: %v",err)
		}

		if response.Result != test.expected {
			t.Errorf("Expected %v, got %v",test.expected, response.Result)
		}
	})
}

}

// def func to test the failures
func TestMathHandlerErrors(t *testing.T){
	tests := []struct {
		name string
		a 	string
		b	string
		operation string
		expectedError string
	}{
		{
			name:	"invalid number",
			a:		"Hello",
			b:		"5",
			operation: "add",
			expectedError: "Please provide a valid number for a and b",
		},
		{
			name:	"invalid operation",
			a:		"10",
			b:	    "5",
			operation: "power",
			expectedError: "Unknown operation. Use 'add', 'sub', 'mul', 'div'",
		},
		{
			name:	"division by zero",
			a:		"10",
			b:		"0",
			operation: "div",
			expectedError: "Cannot divide the number by Zero",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			url := fmt.Sprintf("/calculate?a=%s&b=%s&operation=%s",
		    					test.a,
								test.b,
								test.operation)

			req := httptest.NewRequest(http.MethodGet,url, nil)
			rec := httptest.NewRecorder()

			mathHandler(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf(
					"Expected %d, got %d",
					http.StatusBadRequest, 
					rec.Code,
				)
			}

			var response ErrorResponse
			err := json.NewDecoder(rec.Body).Decode(&response)
			if err != nil {
				t.Fatalf("Failed to decode the response: %v",err)
			}

			if response.Error != test.expectedError {
				t.Errorf(
					"Expected %v, got %v",
					test.expectedError,
					response.Error,
				)
			}


		})
	}
}