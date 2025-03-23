package http

import (
	"log/slog"

	"github.com/NTolka/project-start-of-work/internal/usecase"
	"github.com/gorilla/mux"
)

func NewRouters(usecase *usecase.Usecase, logger *slog.Logger) *mux.Router {
	router := mux.NewRouter()
	handler := NewHandler(usecase, logger)

	router.HandleFunc("/", handler.HandleHello).Methods("GET")

	return router
}
