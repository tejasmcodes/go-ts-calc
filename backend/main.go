package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := http.NewServeMux()

	// Handling entry-point
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Home Page!"))
	})

	// Handling the /hello/name endpoint
	mux.HandleFunc("GET /hello/{name}", helloHandler)

	// Start the server on Post 7070
	fmt.Println("Server is running on http://localhost:7070")
	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Fatalf("Sever failed to start: %v", err)
	}
}
