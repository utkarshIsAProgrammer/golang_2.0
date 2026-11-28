package main

import (
	"log"
	"net/http"
	"time"

	"github.com/utkarshIsAProgrammer/go_tasks/internal/router"
)

func main() {

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router.New(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Println("Server is running on PORT :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
