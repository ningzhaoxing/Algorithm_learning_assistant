package main

import (
	"fmt"
	"getQuestionBot/internal/config"
	"getQuestionBot/internal/service"
	"log"
	"path/filepath"
)

func main() {
	// 加载配置文件
	configPath, err := filepath.Abs("config.yaml")
	if err != nil {
		log.Fatalf("获取配置文件路径失败: %v\n", err)
	}

	// 使用viper加载配置
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v\n", err)
	}

	// 创建服务实例
	dingTalkService := service.NewDingTalkService(cfg)
	leetCodeService := service.NewLeetCodeService(cfg, dingTalkService)

	// 获取并发送力扣信息
	err = leetCodeService.SendWeeklyReport()
	if err != nil {
		fmt.Printf("获取并发送力扣信息失败: %v\n", err)
		return
	}

	// 保持程序运行
	select {}
}
