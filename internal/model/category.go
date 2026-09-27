package model

import (
	"time"
)

// 分类模型
type Category struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Slug        string    `gorm:"uniqueIndex;not null;size:100" json:"slug"`
	Description string    `gorm:"size:255" json:"description"`
	Sort        int       `gorm:"default:0" json:"sort"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
