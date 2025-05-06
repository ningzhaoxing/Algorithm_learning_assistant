package models

// UserWebsite 用户与刷题网站的关联表
type UserWebsite struct {
	UserID    uint    `gorm:"primaryKey"`
	WebsiteID uint    `gorm:"primaryKey"`
	Username  string  `gorm:"not null"`
	UserURL   string  `gorm:"not null"`                           // 用户在该网站的个人首页URL
	User      User    `gorm:"foreignKey:UserID;references:ID"`    // 属于 User
	Website   Website `gorm:"foreignKey:WebsiteID;references:ID"` // 属于 Website
}
