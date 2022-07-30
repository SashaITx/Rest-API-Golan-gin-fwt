package service

import (
	rest "Rest_API_Golan-gin-fwt"
	"Rest_API_Golan-gin-fwt/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
