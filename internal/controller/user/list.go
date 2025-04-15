package user

import (
	"getQuestionBot/internal/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EndpointCtl struct {
	UserRepo dao.UserRepo
}

func NewUserController(userRepo dao.UserRepo) *EndpointCtl {
	return &EndpointCtl{UserRepo: userRepo}
}

func (c *EndpointCtl) ListUsers(ctx *gin.Context) {
	// 获取部门参数，默认为空字符串获取所有用户
	dep := ctx.Query("department")
	// 调用UserRepo接口获取用户数据
	users, err := c.UserRepo.GetUsersByDepartment(dep)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "获取用户数据失败",
		})
		return
	}

	// 渲染模板
	ctx.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}
