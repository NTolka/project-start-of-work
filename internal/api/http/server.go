package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/NTolka/project-start-of-work/internal/config"
	"github.com/NTolka/project-start-of-work/internal/logger"
	"github.com/NTolka/project-start-of-work/internal/repository"
	"github.com/NTolka/project-start-of-work/internal/usecase"
)

type Server struct {
	cfg    *config.Config
	log    *logger.Logger
	repo   *repository.Repository
	server *http.Server
}

func NewServer(cfg *config.Config, log *logger.Logger, repo *repository.Repository) *Server {
	// Инициализация usecase
	service := usecase.NewService(repo, log)

	// Инициализация роутера
	router := NewRouters(repo, service, log.Logger)

	return &Server{
		cfg: cfg,
		log: log,
		server: &http.Server{
			Addr:         ":" + strconv.Itoa(cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
			IdleTimeout:  cfg.Server.IdleTimeout,
			Handler:      router,
		},
	}
}

func (s *Server) Start() error {
	s.log.Logger.Info("Сервер запущен успешно", "port", s.cfg.Server.Port)
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
