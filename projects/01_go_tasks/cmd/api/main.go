package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	// create context that cancels on SIGNINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// buffered channel that hold one error value
	serverErr := make(chan error, 1)

	// start server in a goroutine
	go func() {
		log.Println("Server is running on PORT :8080")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Http server error: %v", err)
			serverErr <- err
		}
	}()

	// wait for shutdown signal
	<-ctx.Done()
	log.Println("Shutdown signal received")

	// gracefully shutdown with a timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown %v", err)
	}

	log.Println("Server exited cleanly")
}
