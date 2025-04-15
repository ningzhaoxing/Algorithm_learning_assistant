package dao

import "getQuestionBot/internal/models"

type UserRepo interface {
	// GetUsersByDepartment 根据部门查询所有用户(包括用户关联表)
	GetUsersByDepartment(dep string) ([]models.User, error)
	// GetUserAndWebsitesByDepartment 根据部门查询所有用户及其刷题网站
	GetUserAndWebsitesByDepartment(dep string, websiteName string) ([]models.UserWebsite, error)
	SaveProblem(problems []models.Problem, uid uint) error
}

type SystemRepo interface {
	// GetSystemConfigById 根据id查询系统信息
	GetSystemConfigById(id uint) (models.System, error)
}
