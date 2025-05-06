package user

import (
	"getQuestionBot/internal/application/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowRegisterPage 展示注册页面
func (c *EndpointCtl) ShowRegisterPage(ctx *gin.Context) {
	// 获取所有网站列表
	websites, err := c.WebsiteRepo.GetAllWebsites()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "获取网站列表失败",
		})
		return
	}

	// 渲染注册页面
	ctx.HTML(http.StatusOK, "register.html", gin.H{
		"websites": websites,
	})
}

// Register 处理用户注册请求
func (c *EndpointCtl) Register(ctx *gin.Context) {
	// 解析请求数据
	var req user.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "请求数据格式错误",
		})
		return
	}

	// 调用注册服务
	if err := c.UserService.Register(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "注册失败: " + err.Error(),
		})
		return
	}

	// 注册成功
	ctx.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}
