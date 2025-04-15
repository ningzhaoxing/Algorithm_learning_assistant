package crawl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ServiceImpl struct {
}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

// GetPageSource 通过url获取静态页面
func (s *ServiceImpl) GetPageSource(url string) ([]byte, error) {
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
	return body, nil
}
