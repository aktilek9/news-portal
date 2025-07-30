package dto

type News struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID int    `json:"author_id"`
}
