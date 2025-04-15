package main

import (
	"getQuestionBot/internal/application"
	"getQuestionBot/internal/config"
	systemRepo "getQuestionBot/internal/dao/system"
	userRepo "getQuestionBot/internal/dao/user"
	"getQuestionBot/internal/models"
	"getQuestionBot/internal/service/impl/crawl"
	"getQuestionBot/internal/service/impl/dingtalk"
	"getQuestionBot/internal/service/impl/message"
	"getQuestionBot/pkg/initizle"
	"log"
)

func main() {
	// 使用viper加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v\n", err)
	}

	db := initizle.DBInit(*cfg)

	err = db.AutoMigrate(
		&models.User{},
		&models.Problem{},
		&models.System{},
		&models.Website{},
		&models.UserWebsite{},
	)

	// 依赖注入
	userRepo := userRepo.NewRepositoryImpl(db)
	sysRepo := systemRepo.NewRepositoryImpl(db)

	crawlImpl := crawl.NewServiceImpl()
	dingtalkImpl := dingtalk.NewServiceImpl(cfg)
	messageImpl := message.NewServiceImpl()

	apply := application.NewService(userRepo, sysRepo, crawlImpl, dingtalkImpl, messageImpl)
	apply.Apply("家族六期", "力扣")
	// 保持程序运行
	//select {}
}
