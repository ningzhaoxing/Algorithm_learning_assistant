package service

import (
	"getQuestionBot/internal/models"
)

type CrawlService interface {
	// GetPageSource 获取静态页面数据
	GetPageSource(url string) ([]byte, error)
}

type MessageProcessService interface {
	// GetProblemListByPageSource 通过静态资源获取题目列表数据
	GetProblemListByPageSource(body []byte) (*models.User, error)
	// MessageAssembly 自定义消息数据组装
	MessageAssembly(users []models.User, system models.System) (string, error)
}

type DingtalkService interface {
	// SendMessage 将消息推送到钉钉机器人
	SendMessage(message string) error
}
