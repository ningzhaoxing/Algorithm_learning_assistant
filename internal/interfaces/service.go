package interfaces

import "getQuestionBot/internal/models"

// LeetCodeService 定义LeetCode服务接口
type LeetCodeService interface {
	// GetUserProfile 获取用户主页信息
	GetUserProfile(url string) (*models.LeetCodeProfile, error)
	// SendWeeklyReport 发送周报
	SendWeeklyReport() error
}

// MessageService 定义消息服务接口
type MessageService interface {
	// SendMessage 发送消息
	SendMessage(message string) error
}