package dao

import "getQuestionBot/internal/models"

type UserRepo interface {
	// GetUsersByDepartment 根据部门查询所有用户(包括用户关联表)
	GetUsersByDepartment(dep string) ([]models.User, error)
	// GetUserAndWebsitesByDepartment 根据部门查询所有用户及其刷题网站
	GetUserAndWebsitesByDepartment(dep string, websiteName string) ([]models.UserWebsite, error)
	// AddUser 新增用户
	AddUser(user models.User) (uint, error)
	// AddUserWebsite 新增用户关联网站
	AddUserWebsite(userWebsite models.UserWebsite) error
}

type ProblemRepo interface {
	// SaveProblem 将获取的题目保存到数据库
	SaveProblem(problems []models.Problem, uid uint) error
}

type SystemRepo interface {
	// GetSystemConfigById 根据id查询系统信息
	GetSystemConfigById(id uint) (models.System, error)
}

type WebsiteRepo interface {
	GetAllWebsites() ([]models.Website, error)
}
