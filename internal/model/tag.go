package model

import (
	"time"
)

// 标签模型
type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Slug      string    `gorm:"uniqueIndex;not null;size:100" json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	Articles  []Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
}

func (Tag) TableName() string {
	return "tags"
}
