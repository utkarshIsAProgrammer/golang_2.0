package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// structure for health response
type HealthResponse struct {
	Status     string `json:"status"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	RemoteAddr string `json:"remoteAddr"`
}

type AboutResponse struct {
	Application string `json:"application"`
	Version     string `json:"version"`
}

// health response handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status:     "ok",
		Method:     r.Method,
		Path:       r.URL.Path,
		RemoteAddr: r.RemoteAddr,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode health response:", err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	respsone := AboutResponse{
		Application: "Go Trello",
		Version:     "1.0.0",
	}

	if err := json.NewEncoder(w).Encode(respsone); err != nil {
		log.Println("failed to encode the about response:", err)
	}

}

func main() {
	// create a router
	mux := http.NewServeMux()

	// create a route
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/about", aboutHandler)

	// server configuration
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server is running on localhost:8080")

	// listen for incoming connection and check for error
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err) //
	}

}
