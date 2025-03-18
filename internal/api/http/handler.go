package http

import (
	"log/slog"
	"net/http"

	"github.com/NTolka/project-start-of-work/internal/usecase"
)

type Handler struct {
	service *usecase.Service
	logger  *slog.Logger
}

func NewHandler(service *usecase.Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
