package main

import (
	"blog-system/internal/handler"
	"blog-system/internal/repository"
	"blog-system/pkg/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	dsn := repository.BuildDSN((cfg.Database))
	db, err := repository.InitDB(dsn)
	if err != nil {
		log.Fatal("初始数据库失败：", err)
	}

	gin.SetMode(cfg.Server.Mode)

	r := gin.Default()
	setupRoutes(r, db)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("服务器启动在 http://%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("启动服务器失败：%v", err)
	}
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	//初始化任务
	userRepo := repository.NewUseRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//初始化jwt
	suth.InitJWT(cfg.JWT.Secret)

	//公开路由
	api := r.Group("api"){
		api.POST("/users/register",userHandler.Register)
		api.POST("/users/login",userHandler.Login)
	}

	//需要认证的路由
	authGroup := api.Group("")
	authGroup.Use(auth.AuthMiddleware())
	{
		authGroup.GET("/users/profile",userHandler.GetProfile)
	}
}
