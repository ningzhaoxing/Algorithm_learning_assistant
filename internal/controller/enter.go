package controller

import (
	"getQuestionBot/internal/application/user"
	userCtl "getQuestionBot/internal/controller/user"
	"getQuestionBot/internal/dao"
)

type apiGroup struct {
	UserSrv userCtl.EndpointCtl
}

var APIs *apiGroup

func InitSrbInject(userRepo dao.UserRepo, userService user.Service, websiteRepo dao.WebsiteRepo) {
	APIs = &apiGroup{
		UserSrv: *userCtl.NewUserController(userRepo, websiteRepo, userService),
	}
}
