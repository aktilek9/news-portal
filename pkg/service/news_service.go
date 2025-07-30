package service

import (
	"news-portal/dto"
	"news-portal/models"
)

func (s *service) GetAllNews() ([]models.News, error) {
	return s.repo.GetAllNews()
}

func (s *service) GetNewsByID(id int) (models.News, error) {
	return s.repo.GetNewsByID(id)
}

func (s *service) CreateNews(newsDTO *dto.News) error {
	news := models.News{
		Title: newsDTO.Title,
		Content: newsDTO.Content,
		AuthorID: uint(newsDTO.AuthorID),
	}
	return s.repo.CreateNews(&news)
}

func (s *service) UpdateNews(id int, news dto.News) error {
	return s.repo.UpdateNews(id, news)
}

func (s *service) DeleteNews(id int) error {
	return s.repo.DeleteNews(id)
}
