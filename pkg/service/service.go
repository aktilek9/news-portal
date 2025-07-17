package service

import "news-portal/pkg/repository"

type service struct {
	repo repository.Repository
}

type Service interface {
}

func NewService(repo repository.Repository) *service {
	return &service{repo: repo}
}
