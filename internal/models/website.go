package models

import "gorm.io/gorm"

// Website 刷题网站表
type Website struct {
	gorm.Model
	Name string `gorm:"not null"` // 网站名称
	URL  string `gorm:"not null"` // 网站URL
}
