package service

import (
	"news-portal/dto"
	"news-portal/pkg/jwt"
	"news-portal/pkg/repository"
)

type service struct {
	repo repository.Repository
	jwt jwt.JWTService
}

type Service interface {
	Login(email, password string) (string, error)
	Register(user *dto.UserDto) (int, error)
}

func NewService(repo repository.Repository, jwt jwt.JWTService) Service {
	return &service{repo: repo, jwt: jwt}
}


