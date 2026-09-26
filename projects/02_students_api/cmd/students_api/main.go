package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/utkarshIsAProgrammer/students_api/config"
)

/* todos (
 * load config
 * database setup
 * setup router
 * setup server
) */

func main() {
	// load config
	cfg := config.MustLoad()

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the student api homepage!"))
	})

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("Server is running on Address %s", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start the server! %s", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down the server!")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown the server!", slog.String("error:", err.Error()))
	}

	slog.Info("Server shutdown successfully!")
}
