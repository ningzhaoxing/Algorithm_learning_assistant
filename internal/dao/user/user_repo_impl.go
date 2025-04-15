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
	if err := r.db.Where("department = ?", dep).
		Preload("Websites").
		Preload("Problems").
		Find(&users).Error; err != nil {
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

// SaveProblem 将为id用户第term学期，第week周的解题列表存入到数据库
func (r *RepositoryImpl) SaveProblem(problems []models.Problem, uid uint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, p := range problems {
			fmt.Println(p)
			p.UserID = uid
			if err := r.db.Model(&models.Problem{}).Create(&p).Error; err != nil {
				fmt.Println(err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
