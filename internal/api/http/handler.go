package http

import (
	"log/slog"
	"net/http"

	"github.com/NTolka/project-start-of-work/internal/repository"
	"github.com/NTolka/project-start-of-work/internal/usecase"
)

type Handler struct {
	repo    *repository.Repository
	usecase *usecase.Usecase
	logger  *slog.Logger
}

func NewHandler(usecase *usecase.Usecase, logger *slog.Logger) *Handler {
	return &Handler{
		usecase: usecase,
		logger:  logger,
	}
}

func (h *Handler) HandleHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
