package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email     string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password  string         `gorm:"not null;size:255" json:"password"`
	Nickname  string         `gorm:"size:50" json:"nickname"`
	Avatar    string         `gorm:"size:255" json:"avatar"`
	Role      string         `gorm:"default:user;size:20" json:"role"`
	Status    string         `gorm:"default:active;size:20" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	//关联
	Articles []Article `gorm:"foreignKey:UserID" json:"articles,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
}

// 指定表名
func (User) TableName() string {
	return "users"
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) IsActive() bool {
	return u.Status == "active"
}
