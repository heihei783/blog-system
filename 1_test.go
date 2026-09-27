package blogsystem

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"blog-system/pkg/config"
	"fmt"
	"log"
	"testing"
)

func TestUserModel(t *testing.T) {
	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	dsn := repository.BuildDSN((cfg.Database))
	db, err := repository.InitDB(dsn)
	if err != nil {
		log.Fatal("初始数据库失败：", err)
	}

	// 测试创建
	if err := db.Create(user).Error; err != nil {
		t.Fatal(err)
	}

	// 测试查询
	var foundUser model.User
	db.First(&foundUser, user.ID)

	if foundUser.Username != user.Username {
		t.Errorf("期望用户名 %s, 得到 %s", user.Username, foundUser.Username)
	}

	fmt.Println("用户创建成功")
}
