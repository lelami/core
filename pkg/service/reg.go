package service

import (
	"core"
	"core/pkg/repository"
)

type RegService struct {
	repo repository.Registration
}

func NewRegService(repo repository.Registration) *RegService {
	return &RegService{repo: repo}
}

func (s *RegService) CreateUser(user core.User) (int, error) {
	return s.repo.CreateUser(user)
}
