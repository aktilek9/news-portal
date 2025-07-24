package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
