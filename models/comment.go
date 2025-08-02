package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
	NewsID   uint   `json:"news_id"`
}

func (Comment) TableName() string {
	return "comment" // Теперь GORM будет использовать `comment`, а не `comments`
}
