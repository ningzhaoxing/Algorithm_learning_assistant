package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"getQuestionBot/internal/config"
	"getQuestionBot/internal/interfaces"
	"net/http"
	"time"
)

// DingTalkServiceImpl 实现MessageService接口
type DingTalkServiceImpl struct {
	config *config.Config
}

// NewDingTalkService 创建钉钉消息服务实例
func NewDingTalkService(config *config.Config) interfaces.MessageService {
	return &DingTalkServiceImpl{
		config: config,
	}
}

// DingTalkMessage 钉钉消息格式
type DingTalkMessage struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// calculateSignature 计算钉钉消息签名
func calculateSignature(secret string, timestamp int64) string {
	strToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(strToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// SendMessage 发送消息
func (s *DingTalkServiceImpl) SendMessage(message string) error {
	fmt.Println(message)
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
	sign := calculateSignature(s.config.DingTalk.Secret, timestamp)

	webhookURL := fmt.Sprintf("%s&timestamp=%d&sign=%s", s.config.DingTalk.Webhook, timestamp, sign)

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
