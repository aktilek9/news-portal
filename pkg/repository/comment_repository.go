package repository

import "news-portal/models"

func (r *repository) CreateComment(comment *models.Comment) error {
	if err := r.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetCommentsByNewsID(newsID int) ([]models.Comment, error) {
	var comment []models.Comment
	if err := r.db.Where("news_id = ?", newsID).Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *repository) GetCommentByID(id int) (models.Comment, error) {
	var comment models.Comment
	if err := r.db.Where("id = ?", id).First(&comment).Error; err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (r *repository) DeleteComment(id int) error {
	result := r.db.Delete(&models.Comment{}, id)
	return result.Error
}
