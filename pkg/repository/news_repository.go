package repository

import (
	"news-portal/dto"
	"news-portal/models"
	"time"
)

func (r *repository) GetAllNews() ([]models.News, error) {
	var news []models.News
	if err := r.db.Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *repository) GetNewsByID(id int) (models.News, error) {
	var news models.News
	result := r.db.First(&news, id)
	if result.Error != nil {
		return models.News{}, result.Error
	}
	return news, nil
}

func (r *repository) CreateNews(news *models.News) error {
	if err := r.db.Create(news).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateNews(id int, newsUpdate dto.News) error {
	result := r.db.Model(&models.News{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":      newsUpdate.Title,
		"content":    newsUpdate.Content,
		"updated_at": time.Now(),
	})
	return result.Error
}

func (r *repository) DeleteNews(id int) error {
	result := r.db.Delete(&models.News{}, id)
	return result.Error
}
