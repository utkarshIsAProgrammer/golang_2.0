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

	// Create and configure the HTTP server.
	server := &http.Server{
		Addr:              ":8080",          // Address and port on which the server listens.
		Handler:           router.New(),     // Register and use our application routes.
		ReadHeaderTimeout: 5 * time.Second,  // Maximum time allowed to read request headers.
		ReadTimeout:       15 * time.Second, // Maximum time allowed to read the entire request.
		WriteTimeout:      30 * time.Second, // Maximum time allowed to write a response.
		IdleTimeout:       60 * time.Second, // Maximum time to wait for another request on an idle connection.
	}

	// Create a context that gets cancelled when the application receives
	// SIGINT (Ctrl+C) or SIGTERM (termination signal).
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop() // Stop signal notification when main exits.

	// Create a buffered channel that can hold one server error.
	// It allows the server goroutine to report unexpected errors.
	serverErr := make(chan error, 1)

	// Start the HTTP server in a separate goroutine because ListenAndServe()
	// blocks while the server is running.
	go func() {
		log.Println("Server is running on PORT :8080")

		// Start accepting HTTP requests.
		// http.ErrServerClosed is expected when the server shuts down normally.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)

			// Send unexpected server errors to the main goroutine.
			serverErr <- err
		}
	}()

	// Wait until the shutdown signal is received.
	// NOTE: This currently waits only for the signal, not serverErr.
	<-ctx.Done()
	log.Println("Shutdown signal received")

	// Create a fresh context with a 10-second timeout for graceful shutdown.
	// We use a fresh context because the signal context above is already cancelled.
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel() // Release resources associated with the timeout context.

	// Gracefully shut down the HTTP server.
	// Stop accepting new connections and allow active requests to finish.
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Log successful completion of graceful shutdown.
	log.Println("Server exited cleanly")
}
