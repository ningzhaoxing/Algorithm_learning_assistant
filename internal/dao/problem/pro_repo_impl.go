package problem

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

// SaveProblem 将为id用户第term学期，第week周的解题列表存入到数据库
func (r *RepositoryImpl) SaveProblem(problems []models.Problem, uid uint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		for _, p := range problems {
			p.UserID = uid
			var existing models.Problem
			if err := r.db.Model(&models.Problem{}).Where("question_id = ? AND user_id = ?", p.QuestionId, p.UserID).First(&existing).Error; err == nil {
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
