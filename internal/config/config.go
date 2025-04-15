package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var path = "./config.yaml"

// Config 应用配置结构
type Config struct {
	DingTalk    DingTalkConfig    `mapstructure:"dingtalk"`
	LeetCode    LeetCodeConfig    `mapstructure:"message"`
	DingMessage DingMessageConfig `mapstructure:"ding_message"`
	Database    DatabaseConfig    `mapstructure:"database"`
}

// DingMessageConfig 钉钉消息配置
type DingMessageConfig struct {
	Header string `mapstructure:"header"`
	Bottom string `mapstructure:"bottom"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// LeetCodeConfig LeetCode配置
type LeetCodeConfig struct {
	Users         []UserConfig `mapstructure:"user"`
	SemesterStart string       `mapstructure:"semester_start"`
	MinimumSolved int          `mapstructure:"minimum_solved"` // 最小解题数量要求
}

// UserConfig 用户配置
type UserConfig struct {
	URL  string `mapstructure:"url"`
	Name string `mapstructure:"name"`
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &config, nil
}
