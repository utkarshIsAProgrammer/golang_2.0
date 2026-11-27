package handler

import (
	"encoding/json"
	"net/http"
)

type AboutResponse struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := AboutResponse{
		Name:        "Go todo api",
		Version:     "1.0.0",
		Description: "A REST API for managing tasks",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
