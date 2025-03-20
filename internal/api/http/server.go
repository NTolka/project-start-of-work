package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/NTolka/project-start-of-work/internal/config"
	"github.com/NTolka/project-start-of-work/internal/logger"
	"github.com/NTolka/project-start-of-work/internal/usecase"

	"github.com/NTolka/project-start-of-work/internal/repository"
)

type Server struct {
	cfg    *config.Config
	log    *logger.Logger
	server *http.Server
}

func NewServer(cfg *config.Config, log *logger.Logger) *Server {
	// Инициализация репозитория и usecase
	repo := repository.NewRepository()
	service := usecase.NewService(repo, log)

	// Инициализация роутера
	router := NewRouters(service, log.Logger)

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
	s.log.Info("Сервер запущен успешно", "port", s.cfg.Server.Port)
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
