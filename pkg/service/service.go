package service

import (
	"core"
	"core/pkg/repository"
)

type Authorization interface {
	CreateAuthUser(user core.AuthUser) (int, error)
	GenerateToken(phone int64, code int) (string, error)
	ParseToken(token string) (int, error)
}

type Registration interface {
	CreateUser(user core.User) (int, error)
}

type Logging interface {
	WriteLog(log string) error
}

type Service struct {
	Registration
	Authorization
	Logging
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Registration:  NewRegService(repos.Registration),
		Authorization: NewAuthService(repos.Authorization),
		Logging:       NewLogService(repos.Logging),
	}
}
