package repository

import (
	"blog-system/internal/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.UserProfile{},
		&model.Category{},
		&model.Tag{},
		&model.Article{},
		&model.ArticleTag{},
		&model.Comment{},
	)
}
