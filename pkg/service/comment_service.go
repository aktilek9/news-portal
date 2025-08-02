package service

import (
	"news-portal/dto"
	"news-portal/models"
)

func (s *service) CreateComment(commentDTO dto.CommentDTO) error {
	comment := models.Comment{
		Content:  commentDTO.Content,
		AuthorID: uint(commentDTO.AuthorID),
		NewsID:   uint(commentDTO.NewsID),
	}
	return s.repo.CreateComment(&comment)
}

func (s *service) GetCommentsByNewsID(newsID int) ([]models.Comment, error) {
	return s.repo.GetCommentsByNewsID(newsID)
}

func (s *service) GetCommentByID(id int) (models.Comment, error) {
	return s.repo.GetCommentByID(id)
}

func (s *service) DeleteComment(id int) error {
	return s.repo.DeleteComment(id)
}
