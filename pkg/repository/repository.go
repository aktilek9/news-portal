package repository

import (
	"news-portal/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateUser(user *models.User) (int, error)
	GetUserByEmail(email string) (*models.User, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
