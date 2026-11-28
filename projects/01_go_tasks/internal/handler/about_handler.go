package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type AboutResponse struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := AboutResponse{
		Name:        "Go tasks",
		Version:     "1.0.0",
		Description: "A REST Api for managing tasks",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal(err, "Error in the about handler")
		return
	}
}
