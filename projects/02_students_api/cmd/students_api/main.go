package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/utkarshIsAProgrammer/students_api/config"
	student "github.com/utkarshIsAProgrammer/students_api/internal/http/handlers/students"
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
	router.HandleFunc("GET /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("Server is running on", slog.String("Address", cfg.HTTPServer.Addr))

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
