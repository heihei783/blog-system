package model

type ArticleTag struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	ArticleID uint `gorm:"index" json:"article_id"`
	TagID     uint `gorm:"index" json:"tag_id"`

	Article Article `gorm:"foreignKey:ArticleID" json:"article,omitempty"`
	Tag     Tag     `gorm:"foreignKey:TagID" json:"tag,omitempty"`
}

func (ArticleTag) TableName() string { return "article_tags" }
