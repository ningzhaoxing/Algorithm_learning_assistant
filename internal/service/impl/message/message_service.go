package message

import (
	"encoding/json"
	"fmt"
	"getQuestionBot/internal/models"
	"getQuestionBot/pkg/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ServiceImpl struct {
}

// NewServiceImpl 创建服务实例
func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

// GetProblemListByPageSource 从获取的静态资源中解析数据
func (s *ServiceImpl) GetProblemListByPageSource(body []byte) (*models.User, error) {
	// 解析响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v, 响应内容: %s", err, string(body))
	}

	// 检查错误
	if errors, ok := result["errors"].([]interface{}); ok && len(errors) > 0 {
		return nil, fmt.Errorf("API返回错误: %v", errors)
	}

	// 提取数据
	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("响应格式错误: %s", string(body))
	}

	submissions, ok := data["recentACSubmissions"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("未找到提交记录: %s", string(body))
	}

	// 构造返回结果
	var problems []models.Problem

	// 获取提交记录
	oneWeekAgo := time.Now().AddDate(0, 0, -7)
	// 使用map来存储已处理的题目ID
	processedQuestions := make(map[string]models.Problem)

	for _, sub := range submissions {
		subMap := sub.(map[string]interface{})
		question := subMap["question"].(map[string]interface{})
		submitTime := int64(subMap["submitTime"].(float64))
		submitTimeObj := time.Unix(submitTime, 0)

		// 只处理一周内的题目
		if submitTimeObj.After(oneWeekAgo) {
			questionId := question["questionFrontendId"].(string)
			// 如果题目ID已存在，只更新最新的提交时间
			if _, exists := processedQuestions[questionId]; !exists {
				processedQuestions[questionId] = models.Problem{
					Number:          fmt.Sprintf("%v", subMap["submissionId"]),
					Title:           question["title"].(string),
					TranslatedTitle: question["translatedTitle"].(string),
					TitleSlug:       question["titleSlug"].(string),
					QuestionId:      questionId,
					SubmitTime:      submitTimeObj.Format("2006-01-02 15:04:05"),
				}
			}
		}
	}

	// 将map转换为切片
	for _, problem := range processedQuestions {
		problems = append(problems, problem)
	}

	// 按提交时间倒序排序
	sort.Slice(problems, func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02 15:04:05", problems[i].SubmitTime)
		tj, _ := time.Parse("2006-01-02 15:04:05", problems[j].SubmitTime)
		return ti.After(tj)
	})

	solvedCount := fmt.Sprintf("%d", len(problems))

	solvedNum, _ := strconv.Atoi(solvedCount)
	profile := &models.User{
		Problems:  problems,
		SolvedNum: solvedNum,
	}

	return profile, nil
}

// MessageAssembly 消息组装
func (s *ServiceImpl) MessageAssembly(users []models.User, system models.System) (string, error) {
	// 计算当前是第几周
	weekNumber := utils.CalCurWeek(system)

	var message strings.Builder
	message.WriteString(fmt.Sprintf("💌【力扣刷题周报·第%d周】💌\n", weekNumber))

	// 添加消息头部
	message.WriteString(system.DingHeader)

	// 按解题数量对用户进行排序
	sort.Slice(users, func(i, j int) bool {
		return users[i].SolvedNum > users[j].SolvedNum
	})

	// 展示前四名用户
	message.WriteString("🌟 本周优秀：\n")
	for i := 0; i < len(users) && i < 4; i++ {
		message.WriteString(fmt.Sprintf("%s ",
			users[i].Name))
	}
	message.WriteString("\n")

	// 将用户按解题数量分组
	var belowMinimum []models.User
	var aboveMinimum []models.User

	// 分组处理
	for _, user := range users {
		if user.SolvedNum < system.MinimumSolved {
			belowMinimum = append(belowMinimum, user)
		} else {
			aboveMinimum = append(aboveMinimum, user)
		}
	}

	// 展示达标的用户
	if len(aboveMinimum) > 0 {
		message.WriteString("✨达标进度同学：\n")
		for _, profile := range aboveMinimum {
			message.WriteString(fmt.Sprintf("%s (本周解题数量：%d)\n", profile.Name, profile.SolvedNum))
		}
	}

	// 如果有未达标的用户，单独展示
	if len(belowMinimum) > 0 {
		message.WriteString(fmt.Sprintf("⚠️ 未达到最低解题要求（%d题）的同学：\n", system.MinimumSolved))
		for _, profile := range belowMinimum {
			message.WriteString(fmt.Sprintf("%s (本周解题数量：%d)\n", profile.Name, profile.SolvedNum))
		}
	}

	// 添加消息底部
	message.WriteString(system.DingBottom)

	msg := strings.ReplaceAll(message.String(), "\\n", "\n")
	return msg, nil
}
