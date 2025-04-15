package config

// DingTalkConfig 钉钉机器人配置
type DingTalkConfig struct {
	Webhook string // Webhook地址
	Secret  string // 加签密钥
}

// NewDingTalkConfig 创建钉钉机器人配置
func NewDingTalkConfig(webhook, secret string) *DingTalkConfig {
	return &DingTalkConfig{
		Webhook: webhook,
		Secret:  secret,
	}
}
