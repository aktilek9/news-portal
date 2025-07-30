package repository

import (
	"news-portal/dto"
	"news-portal/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateUser(user *models.User) (int, error)
	GetUserByEmail(email string) (*models.User, error)
	GetAllNews() ([]models.News, error)
	GetNewsByID(id int) (models.News, error)
	CreateNews(news *models.News) error
	UpdateNews(id int, news dto.News) error
	DeleteNews(id int) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
