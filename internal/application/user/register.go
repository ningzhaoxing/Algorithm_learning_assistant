package user

import (
	"getQuestionBot/internal/dao"
	"getQuestionBot/internal/models"
)

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	WebsiteID  uint   `json:"websiteId"`
	Username   string `json:"username"`
	UserURL    string `json:"userUrl"`
}

type Register struct {
	dao.UserRepo
}

func NewRegister(userRepo dao.UserRepo) *Register {
	return &Register{
		UserRepo: userRepo,
	}
}

// Register 用户注册
func (s *Register) Register(req RegisterRequest) error {
	// 1. 创建用户
	user := models.User{
		Name:       req.Name,
		Department: req.Department,
	}

	// 添加用户到数据库
	id, err := s.UserRepo.AddUser(user)
	if err != nil {
		return err
	}

	// 2. 创建用户与网站的关联
	userWebsite := models.UserWebsite{
		UserID:    id,
		WebsiteID: req.WebsiteID,
		Username:  req.Username,
		UserURL:   req.UserURL,
	}

	// 添加关联关系到数据库
	if err := s.UserRepo.AddUserWebsite(userWebsite); err != nil {
		return err
	}

	return nil
}
