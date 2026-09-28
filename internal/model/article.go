package model

import (
	"time"

	"gorm.io/gorm"
)

// Article 文章模型
type Article struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	Title        string     `gorm:"not null;size:200;index" json:"title"`
	Slug         string     `gorm:"uniqueIndex;not null;size:255" json:"slug"`
	Content      string     `gorm:"type:text" json:"content"`
	Summary      string     `gorm:"size:500" json:"summary"`
	CoverImage   string     `gorm:"size:255" json:"cover_image"`
	Status       string     `gorm:"default:draft;size:20" json:"status"`
	ViewCount    int        `gorm:"default:0" json:"view_count"`
	LikeCount    int        `gorm:"default:0" json:"like_count"`
	CommentCount int        `gorm:"default:0" json:"comment_count"`
	PublishedAt  *time.Time `json:"published_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`

	//外键
	UserID     uint `gorm:"not null;index" json:"user_id"`
	CategoryID uint `gorm:"index" json:"category_id"`

	//关联
	User     User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Category Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags     []Tag     `gorm:"many2many:article_tags;" json:"tags,omitempty"`
	Comments []Comment `gorm:"foreignKey:ArticleID" json:"comments,omitempty"`
}

func (Article) TableName() string { return "articles" }

// BeforeCreate 创建前钩子
func (a *Article) BeforeCreate(tx *gorm.DB) error {
	//生成slug
	if a.Slug == "" {
		// a.Slug = generateSlug(a.Title)
	}

	if a.Status == "published" && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}

	return nil
}

// BeforeUpdate 更新前钩子
func (a *Article) BeforeUpdate(tx *gorm.DB) error {

	if a.Status == "published" && a.PublishedAt == nil {
		now := time.Now()
		a.PublishedAt = &now
	}

	return nil
}

func (a *Article) IsPublished() bool {
	return a.Status == "published"
}

func (a *Article) IncrementViewCount() {
	a.ViewCount++
}
