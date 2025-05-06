package user

import (
	"getQuestionBot/internal/application/user"
	"getQuestionBot/internal/dao"
)

type EndpointCtl struct {
	UserRepo    dao.UserRepo
	WebsiteRepo dao.WebsiteRepo
	UserService user.Service
}

func NewUserController(userRepo dao.UserRepo, websiteRepo dao.WebsiteRepo, userService user.Service) *EndpointCtl {
	return &EndpointCtl{
		UserRepo:    userRepo,
		WebsiteRepo: websiteRepo,
		UserService: userService,
	}
}
