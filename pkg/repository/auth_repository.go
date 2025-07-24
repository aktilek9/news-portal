package repository

import "news-portal/models"

func (r *repository) CreateUser(user *models.User) (int, error) {
	if err := r.db.Create(user).Error; err != nil {
		return 0, err
	}
	return int(user.Model.ID), nil
}

func (r *repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
