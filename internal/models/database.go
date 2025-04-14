package models

import "gorm.io/gorm"

// User 用户表
type User struct {
	gorm.Model
	Username string   `gorm:"uniqueIndex;not null"`
	Name     string   `gorm:"not null"`
	Systems  []System `gorm:"many2many:user_system;"`
}

type Problem struct {
	gorm.Model
	Number          string `gorm:"not null"` // 提交ID
	Title           string // 英文标题
	TranslatedTitle string `gorm:"not null"` // 中文标题
	TitleSlug       string // 题目标识
	QuestionId      string // 题目编号
	SubmitTime      string `gorm:"not null"` // 提交时间
}

// System 刷题网站表
type System struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;not null"` // 网站名称
	URL  string `gorm:"not null"`             // 网站URL
}

// UserSystem 用户与刷题网站的关联表
type UserSystem struct {
	UserID   uint   `gorm:"primaryKey"`
	SystemID uint   `gorm:"primaryKey"`
	UserURL  string `gorm:"not null"` // 用户在该网站的个人首页URL
}
