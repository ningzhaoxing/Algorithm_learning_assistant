package main

import (
	"fmt"
	"getQuestionBot/internal/application"
	"getQuestionBot/internal/config"
	"getQuestionBot/internal/controller"
	systemRepo "getQuestionBot/internal/dao/system"
	userRepo "getQuestionBot/internal/dao/user"
	"getQuestionBot/internal/models"
	"getQuestionBot/internal/service/impl/crawl"
	"getQuestionBot/internal/service/impl/dingtalk"
	"getQuestionBot/internal/service/impl/message"
	"getQuestionBot/pkg/initizle"
	"github.com/robfig/cron/v3"
	"log"
	"time"
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

	controller.InitSrbInject(userRepo)

	apply := application.NewService(userRepo, sysRepo, crawlImpl, dingtalkImpl, messageImpl)

	apply.Apply("家族六期", "力扣")
	// 设置定时任务，每周一早上8:30执行
	c := cron.New(cron.WithLocation(time.FixedZone("CST", 8*3600)))
	_, err = c.AddFunc("30 8 * * 1", func() {
		apply.Apply("家族六期", "力扣")
	})
	if err != nil {
		log.Fatalf("定时任务设置失败: %v\n", err)
	}
	c.Start()

	e := initizle.RouterInit()
	err = e.Run(fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port))
	if err != nil {
		return
	}
}
