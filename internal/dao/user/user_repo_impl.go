package user

import (
	"fmt"
	"getQuestionBot/internal/models"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

// GetUsersByDepartment 根据部门查询所有用户(包括用户关联表)
func (r *RepositoryImpl) GetUsersByDepartment(dep string) ([]models.User, error) {
	var users []models.User
	if err := r.db.Model(&models.User{}).Where("department = ?", dep).
		Preload("Problems").
		Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

// GetUserAndWebsitesByDepartment 根据部门查询所有用户以及用户关联name的网站
func (r *RepositoryImpl) GetUserAndWebsitesByDepartment(dep string, websiteName string) ([]models.UserWebsite, error) {
	// 查询id对应的网站
	var id int
	if err := r.db.Model(&models.Website{}).Select("id").Where("name=?", "力扣").First(&id).Error; err != nil {
		return nil, err
	}

	// 查询用户关联网站
	var websites []models.UserWebsite
	if err := r.db.Model(&models.UserWebsite{}).Preload("User").Where("website_id=?", id).Find(&websites).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return websites, nil
}

// AddUser 新增用户
func (r *RepositoryImpl) AddUser(user models.User) (uint, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, fmt.Errorf("添加用户失败: %w", err)
	}
	return user.ID, nil
}

// AddUserWebsite 新增用户关联网站
func (r *RepositoryImpl) AddUserWebsite(userWebsite models.UserWebsite) error {
	if err := r.db.Create(&userWebsite).Error; err != nil {
		return fmt.Errorf("添加用户关联网站失败: %w", err)
	}
	return nil
}
