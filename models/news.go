package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
	Comments []Comment `json:"comments"`
}
