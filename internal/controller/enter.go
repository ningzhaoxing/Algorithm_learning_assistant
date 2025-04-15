package controller

import (
	userCtl "getQuestionBot/internal/controller/user"
	"getQuestionBot/internal/dao"
)

type apiGroup struct {
	UserSrv userCtl.EndpointCtl
}

var APIs *apiGroup

func InitSrbInject(userRepo dao.UserRepo) {
	APIs = &apiGroup{
		UserSrv: userCtl.EndpointCtl{UserRepo: userRepo},
	}
}
