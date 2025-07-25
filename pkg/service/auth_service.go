package service

import (
	"errors"
	"news-portal/dto"
	"news-portal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := s.jwt.GenerateToken(int(user.Model.ID), user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *service) Register(userDto *dto.UserDto) (int, error) {
	_, err := s.repo.GetUserByEmail(userDto.Email)

	if err == nil {
		return 0, errors.New("user with this email already exists")
	}

	passwordHash, err := hashPassword(userDto.Password)
	if err != nil {
		return 0, err
	}

	user := &models.User{
		Role:     "client",
		Email:    userDto.Email,
		Password: passwordHash,
	}

	return s.repo.CreateUser(user)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
