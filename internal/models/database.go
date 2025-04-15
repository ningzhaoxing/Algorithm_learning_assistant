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

// User 用户表
type User struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	Websites     []Website     `gorm:"many2many:user_websites"`
	UserWebsites []UserWebsite `gorm:"foreignKey:UserID"`
	Problems     []Problem     `gorm:"foreignKey:UserID"`
	Department   string        `gorm:""`
	SolvedNum    int           `gorm:"-"`
}

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
}

// Website 刷题网站表
type Website struct {
	gorm.Model
	Name string `gorm:"not null"` // 网站名称
	URL  string `gorm:"not null"` // 网站URL
}

// UserWebsite 用户与刷题网站的关联表
type UserWebsite struct {
	UserID    uint    `gorm:"primaryKey"`
	WebsiteID uint    `gorm:"primaryKey"`
	Username  string  `gorm:"not null"`
	UserURL   string  `gorm:"not null"`                           // 用户在该网站的个人首页URL
	User      User    `gorm:"foreignKey:UserID;references:ID"`    // 属于 User
	Website   Website `gorm:"foreignKey:WebsiteID;references:ID"` // 属于 Website
}
