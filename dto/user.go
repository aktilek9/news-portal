package dto

type UserDto struct {
	Role     string `json:"role"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password"`
}
