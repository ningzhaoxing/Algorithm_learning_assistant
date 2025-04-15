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
	fmt.Println(users[0].Problems[0])
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
			p.UserID = uid
			var existing models.Problem
			if err := r.db.Model(&models.Problem{}).Where("question_id = ?", p.QuestionId).First(&existing).Error; err == nil {
				// 记录已存在，只更新sub_time
				if err := r.db.Model(&models.Problem{}).Where("id = ?", existing.ID).Update("submit_time", p.SubmitTime).Error; err != nil {
					fmt.Println(err)
					return err
				}
			} else {
				// 记录不存在，创建新记录
				if err := r.db.Model(&models.Problem{}).Create(&p).Error; err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
