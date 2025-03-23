package usecase

import (
	"github.com/NTolka/project-start-of-work/internal/logger"
	"github.com/NTolka/project-start-of-work/internal/repository"
)

type Usecase struct {
	repo *repository.Repository
	log  *logger.Logger
}

func NewService(repo *repository.Repository, log *logger.Logger) *Usecase {
	return &Usecase{
		repo: repo,
		log:  log,
	}
}
