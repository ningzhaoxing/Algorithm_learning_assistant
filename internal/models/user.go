package models

import "gorm.io/gorm"

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
