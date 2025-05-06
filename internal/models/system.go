package models

import "gorm.io/gorm"

// System 配置表
type System struct {
	gorm.Model
	MinimumSolved int    `gorm:"not null"`  // 最小解题数量
	SemesterStart string `gorm:"not null"`  // 学期开始时间
	CurTerm       string `gorm:"not null"`  // 当前学期
	DingHeader    string `gorm:"type:text"` // 钉钉消息头部
	DingBottom    string `gorm:"type:text"` // 钉钉消息底部
}
