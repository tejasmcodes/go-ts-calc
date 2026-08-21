package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type MathResponse struct {
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	// let the query parameter be (?a=1&b=2&operation=add)
	strA := r.URL.Query().Get("a")
	strB := r.URL.Query().Get("b")
	op := r.URL.Query().Get("operation")

	numA, errA := strconv.ParseFloat(strA, 64)
	numB, errB := strconv.ParseFloat(strB, 64)

	if errA != nil || errB != nil {
		errorResponse := ErrorResponse{
			Error: "Please provide a valid number for a and b",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	result := 0.0
	switch op {
	case "add":
		result = numA + numB

	case "sub":
		result = numA - numB

	case "mul":
		result = numA * numB

	case "div":
		if numB == 0.0 {
			errorResponse := ErrorResponse{
				Error: "Cannot divide the number by Zero",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		result = numA / numB

	default:
		errorResponse := ErrorResponse{
			Error: "Unknown operation. Use 'add', 'sub', 'mul', 'div'",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := MathResponse{
		A:         numA,
		B:         numB,
		Operation: op,
		Result:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Router initialized
	mux := http.NewServeMux()

	// Handling entry-point
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home Page!"))
	})

	// Handling the /hello/name endpoint
	mux.HandleFunc("GET /hello/{name}", helloHandler)

	mux.HandleFunc("GET /calculate", mathHandler)

	// Start the server on Port 7070
	fmt.Println("Server is running on http://localhost:7070")
	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Fatalf("Sever failed to start: %v", err)
	}
}
