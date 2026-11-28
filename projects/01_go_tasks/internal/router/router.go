package router

import (
	"net/http"

	"github.com/utkarshIsAProgrammer/go_tasks/internal/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthHandler)
	mux.HandleFunc("GET /about", handler.AboutHandler)

	return mux

}
