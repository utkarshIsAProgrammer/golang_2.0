package main

import (
	"log"
	"net/http"

	"github.com/utkarshIsAProgrammer/go_task_manager/internal/router"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.New(),
	}

	log.Println("Server is runnig on PORT :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
