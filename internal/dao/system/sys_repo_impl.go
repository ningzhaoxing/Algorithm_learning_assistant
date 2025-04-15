package system

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

// GetSystemConfigById 根据id查询系统信息
func (r *RepositoryImpl) GetSystemConfigById(id uint) (models.System, error) {
	var system models.System
	result := r.db.Where("id = ?", id).First(&system)
	if result.Error != nil {
		// 如果找不到记录，返回空的系统配置
		return models.System{}, result.Error
	}
	return system, nil
}
