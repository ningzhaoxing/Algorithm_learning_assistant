package initizle

import (
	"getQuestionBot/internal/controller"
	"github.com/gin-gonic/gin"
)

// RouterInit 初始化路由
func RouterInit() *gin.Engine {
	// 创建默认的gin引擎
	r := gin.Default()

	// 配置中间件
	r.Use(gin.Recovery())

	// 配置静态文件服务
	r.Static("/static", "./static")

	// 配置HTML模板
	r.LoadHTMLGlob("view/pages/*")

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		user := api.Group("/user")
		{
			user.GET("/list", controller.APIs.UserSrv.ListUsers)
			user.GET("/register")
			user.POST("/register", controller.APIs.UserSrv.Register)
		}

		// 问题相关路由
		//problem := api.Group("/problem")
		//{
		//	// TODO: 添加问题相关路由处理器
		//}
		//
		//// 系统相关路由
		//system := api.Group("/system")
		//{
		//	// TODO: 添加系统相关路由处理器
		//}
	}

	return r
}
