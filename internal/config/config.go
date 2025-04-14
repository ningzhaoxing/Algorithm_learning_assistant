package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Config 应用配置结构
type Config struct {
	DingTalk DingTalkConfig `mapstructure:"dingtalk"`
	LeetCode LeetCodeConfig `mapstructure:"leetcode"`
	DingMessage DingMessageConfig `mapstructure:"ding_message"`
}

// DingMessageConfig 钉钉消息配置
type DingMessageConfig struct {
	Header string `mapstructure:"header"`
	Bottom string `mapstructure:"bottom"`
}

// LeetCodeConfig LeetCode配置
type LeetCodeConfig struct {
	Users         []UserConfig `mapstructure:"users"`
	SemesterStart string       `mapstructure:"semester_start"`
	MinimumSolved int          `mapstructure:"minimum_solved"` // 最小解题数量要求
}

// UserConfig 用户配置
type UserConfig struct {
	URL  string `mapstructure:"url"`
	Name string `mapstructure:"name"`
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 验证配置
	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// validateConfig 验证配置
func validateConfig(config *Config) error {
	// 验证钉钉配置
	if config.DingTalk.Webhook == "" {
		return fmt.Errorf("钉钉webhook不能为空")
	}
	if config.DingTalk.Secret == "" {
		return fmt.Errorf("钉钉secret不能为空")
	}
	
	// 验证钉钉消息配置
	if config.DingMessage.Header == "" {
		return fmt.Errorf("钉钉消息header不能为空")
	}
	if config.DingMessage.Bottom == "" {
		return fmt.Errorf("钉钉消息bottom不能为空")
	}

	// 验证LeetCode配置
	if len(config.LeetCode.Users) == 0 {
		return fmt.Errorf("LeetCode用户列表不能为空")
	}

	// 验证学期开始时间
	_, err := time.Parse("2006-01-02", config.LeetCode.SemesterStart)
	if err != nil {
		return fmt.Errorf("学期开始时间格式错误: %w", err)
	}

	return nil
}