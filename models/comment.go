package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
	NewsID   uint   `json:"news_id"`
}
