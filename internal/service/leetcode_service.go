package service

import (
	"fmt"
	"getQuestionBot/internal/config"
	"getQuestionBot/internal/interfaces"
	"getQuestionBot/internal/models"
	"sort"
	"strings"
	"time"
)

// LeetCodeServiceImpl 实现LeetCodeService接口
type LeetCodeServiceImpl struct {
	config         *config.Config
	messageService interfaces.MessageService
}

// NewLeetCodeService 创建LeetCode服务实例
func NewLeetCodeService(config *config.Config, messageService interfaces.MessageService) interfaces.LeetCodeService {
	return &LeetCodeServiceImpl{
		config:         config,
		messageService: messageService,
	}
}

// GetUserProfile 获取用户主页信息
func (s *LeetCodeServiceImpl) GetUserProfile(url string) (*models.LeetCodeProfile, error) {
	return GetLeetCodeProfile(url)
}

// SendWeeklyReport 发送周报
func (s *LeetCodeServiceImpl) SendWeeklyReport() error {
	// 计算当前是第几周
	semesterStart, _ := time.Parse("2006-01-02", s.config.LeetCode.SemesterStart)
	currentTime := time.Now()
	daysSinceStart := currentTime.Sub(semesterStart).Hours() / 24
	weekNumber := int(daysSinceStart/7) + 1

	var message strings.Builder
	message.WriteString(fmt.Sprintf("💌【力扣刷题周报·第%d周】💌\n", weekNumber))

	// 添加消息头部
	message.WriteString(s.config.DingMessage.Header)

	// 获取所有用户的解题信息
	var allProfiles []*models.LeetCodeProfile

	// 获取所有用户的解题信息
	for _, user := range s.config.LeetCode.Users {
		profile, err := s.GetUserProfile(user.URL)
		if err != nil {
			fmt.Printf("获取用户信息失败 [%s]: %v\n", user.URL, err)
			continue
		}

		// 将解题数量转换为整数进行比较
		solvedCount := 0
		fmt.Sscanf(profile.SolvedCount, "%d", &solvedCount)
		profile.SolvedCount = fmt.Sprintf("%d", solvedCount)

		allProfiles = append(allProfiles, profile)
	}

	// 按解题数量对用户进行排序
	sort.Slice(allProfiles, func(i, j int) bool {
		var si, sj int
		fmt.Sscanf(allProfiles[i].SolvedCount, "%d", &si)
		fmt.Sscanf(allProfiles[j].SolvedCount, "%d", &sj)
		return si > sj
	})

	// 展示前四名用户
	message.WriteString("🌟 本周优秀：\n")
	for i := 0; i < len(allProfiles) && i < 4; i++ {
		message.WriteString(fmt.Sprintf("%s ",
			allProfiles[i].Username))
	}
	message.WriteString("\n")

	// 将用户按解题数量分组
	var belowMinimum []*models.LeetCodeProfile
	var aboveMinimum []*models.LeetCodeProfile

	// 分组处理
	for _, profile := range allProfiles {
		solvedCount := 0
		fmt.Sscanf(profile.SolvedCount, "%d", &solvedCount)

		if solvedCount < s.config.LeetCode.MinimumSolved {
			belowMinimum = append(belowMinimum, profile)
		} else {
			aboveMinimum = append(aboveMinimum, profile)
		}
	}

	// 展示达标的用户
	if len(aboveMinimum) > 0 {
		message.WriteString("达标进度：\n")
		for _, profile := range aboveMinimum {
			message.WriteString(fmt.Sprintf("%s (本周解题数量：%s)\n", profile.Username, profile.SolvedCount))
		}
	}

	// 如果有未达标的用户，单独展示
	if len(belowMinimum) > 0 {
		message.WriteString(fmt.Sprintf("⚠️ 未达到最低解题要求（%d题）的用户：\n", s.config.LeetCode.MinimumSolved))
		for _, profile := range belowMinimum {
			message.WriteString(fmt.Sprintf("%s (本周解题数量：%s)\n", profile.Username, profile.SolvedCount))
		}
	}

	// 添加消息底部
	message.WriteString(s.config.DingMessage.Bottom)
	return s.messageService.SendMessage(message.String())
}
