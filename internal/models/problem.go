package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Number          string `gorm:"not null"` // 提交ID
	Title           string // 英文标题
	TranslatedTitle string `gorm:"not null"` // 中文标题
	TitleSlug       string // 题目标识
	QuestionId      string // 题目编号
	SubmitTime      string `gorm:"not null"` // 提交时间
	Term            string `gorm:"not null"` // 学期
	Week            string `gorm:"not null"` // 周数
	UserID          uint   `gorm:"index"`    // 外键关联User表
	Url             string `gorm:""`         // 题目url
	Difficulty      string `gorm:""`         // 难度
}
