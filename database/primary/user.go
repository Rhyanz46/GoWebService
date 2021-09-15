package primary

import "time"

type User struct {
	Email      string    `gorm:"unique;primary_key;size:100;index:idx_user_email;" json:"email"`
	FullName   string    `gorm:"size:100;" json:"fullname"`
	Username   string    `gorm:"size:200;index:idx_user_username;" json:"username"`
	Password   string    `gorm:"index:idx_user_key;" json:"-"`
	Role       int       `gorm:"default:0" json:"role"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_time"`
}