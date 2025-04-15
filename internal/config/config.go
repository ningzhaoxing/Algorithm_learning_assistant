package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var path = "./config.yaml"

// Config 应用配置结构
type Config struct {
	DingTalk DingTalkConfig `mapstructure:"dingtalk"`
	Database DatabaseConfig `mapstructure:"database"`
	App      AppConfig      `mapstructure:"app"`
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// UserConfig 用户配置
type UserConfig struct {
	URL  string `mapstructure:"url"`
	Name string `mapstructure:"name"`
}

// AppConfig App配置
type AppConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
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
