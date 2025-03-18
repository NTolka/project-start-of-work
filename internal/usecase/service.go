package usecase

import (
	"github.com/NTolka/project-start-of-work/internal/logger"
	"github.com/NTolka/project-start-of-work/internal/repository"
)

type Service struct {
	repo *repository.Repository
	log  *logger.Logger
}

func NewService(repo *repository.Repository, log *logger.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}
