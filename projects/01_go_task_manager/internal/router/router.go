package router

import (
	"net/http"

	"github.com/utkarshIsAProgrammer/go_task_manager/internal/handler"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/about", handler.AboutHandler)

	return mux
}
