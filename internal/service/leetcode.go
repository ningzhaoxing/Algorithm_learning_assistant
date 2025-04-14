package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"getQuestionBot/internal/config"
	"getQuestionBot/internal/models"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// GetLeetCodeProfile 获取力扣个人主页信息
func GetLeetCodeProfile(url string) (*models.LeetCodeProfile, error) {
	// 从URL中提取用户名
	parts := strings.Split(url, "/")
	username := parts[len(parts)-2]

	// 构造API请求URL
	apiURL := "https://leetcode.cn/graphql/noj-go/"

	// 创建会话
	client := &http.Client{}

	// 设置User-Agent
	userAgent := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"

	// 构造GraphQL查询
	query := map[string]interface{}{
		"query": `query recentACSubmissions($userSlug: String!) {
			recentACSubmissions(userSlug: $userSlug) {
				submissionId
				submitTime
				question {
					title
					translatedTitle
					titleSlug
					questionFrontendId
				}
			}
		}`,
		"variables": map[string]interface{}{
			"userSlug": username,
		},
		"operationName": "recentACSubmissions",
	}

	// 将查询转换为JSON
	payload, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("构造请求失败: %v", err)
	}

	// 创建POST请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://leetcode.cn/accounts/login/")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

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
	profile := models.NewLeetCodeProfile(username, solvedCount, problems)

	return profile, nil
}

// SendToDingTalk 发送消息到钉钉
func SendToDingTalk(config *config.DingTalkConfig, message string) error {
	msg := DingTalkMessage{
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: message,
		},
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("消息序列化失败: %v", err)
	}

	timestamp := time.Now().UnixMilli()
	sign := calculateSignature(config.Secret, timestamp)

	webhookURL := fmt.Sprintf("%s&timestamp=%d&sign=%s", config.Webhook, timestamp, sign)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("发送消息失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("钉钉返回错误状态码: %d", resp.StatusCode)
	}

	return nil
}
