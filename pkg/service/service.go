package service

import (
	"news-portal/dto"
	"news-portal/models"
	"news-portal/pkg/jwt"
	"news-portal/pkg/repository"
)

type service struct {
	repo repository.Repository
	jwt  jwt.JWTService
}

type Service interface {
	Login(email, password string) (string, error)
	Register(user *dto.UserDto) (int, error)
	GetAllNews() ([]models.News, error)
	GetNewsByID(id int) (models.News, error)
	CreateNews(news dto.News) error
	UpdateNews(id int, news dto.News) error
	DeleteNews(id int) error
	CreateComment(commentDTO dto.CommentDTO) error
	GetCommentsByNewsID(newsID int) ([]models.Comment, error)
	GetCommentByID(id int) (models.Comment, error)
	DeleteComment(id int) error
}

func NewService(repo repository.Repository, jwt jwt.JWTService) Service {
	return &service{repo: repo, jwt: jwt}
}
