package http

import (
	"log/slog"

	"github.com/NTolka/project-start-of-work/internal/usecase"
	"github.com/gorilla/mux"
)

func NewRouters(service *usecase.Service, logger *slog.Logger) *mux.Router {
	router := mux.NewRouter()
	handler := NewHandler(service, logger)

	router.HandleFunc("/", handler.HandleHello).Methods("GET")

	return router
}
