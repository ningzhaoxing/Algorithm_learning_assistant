package website

import (
	"getQuestionBot/internal/models"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

// GetAllWebsites 查询当前支持刷题网站列表
func (r *RepositoryImpl) GetAllWebsites() ([]models.Website, error) {
	var websites []models.Website
	if err := r.db.Model(&models.Website{}).Find(&websites).Error; err != nil {
		return nil, err
	}
	return websites, nil
}
