package model

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Status    string    `gorm:"default:pending;size:20" json:"status"`
	IP        string    `gorm:"size:255" json:"-"`
	UserAgent string    `gorm:"size:255" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//外键
	ArticleID uint  `gorm:"not null;index" json:"article_id"`
	UserID    uint  `gorm:"index" json:"user_id"`
	ParentID  *uint `gorm:"index" json:"parent_id,omitempty"`

	//关联
	Article Article   `gorm:"foreignKey:ArticleID" json:"article,omitempty"`
	User    User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Parent  *Comment  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) IsApproved() bool {
	return c.Status == "approved"
}

func (c *Comment) IsReply() bool {
	return c.ParentID != nil
}
