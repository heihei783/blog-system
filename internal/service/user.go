package service

import (
	"blog-system/internal/model"
	"blog-system/internal/repository"
	"blog-system/pkg/auth"
	"blog-system/pkg/utils"
	"errors"
	"time"
)

type UserService interface {
	Register(req RegisterRequest) (*model.User, error)
	Login(req LoginRequest) (string, error)
	GetProfile(userID uint) (*model.User, error)
	UpdateProfile(userID uint, req UpdateProfileRequest) error
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min3,max=50"`
	Email    string `josn:"email" binding:"required,email"`
	Password string `josn:"password" binding:"required,min=6"`
	Nickname string `json:"nickname"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `josn:"avatar"`
}

func (s *UserServiceImpl) Reginster(req RegisterRequest) (*model.User, error) {
	//检查用户名是否存在
	if s.userRepo.ExistsByUsername(req.Username) {
		return nil, errors.New("用户名已存在")
	}

	//检查邮箱是否存在
	if s.userRepo.ExistsByEmail(req.Email) {
		return nil, errors.New("邮箱已存在")
	}

	//加密密码
	hashPassword, err := utils.HashPassword((req.Password))
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	//创建用户
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
		Nickname: req.Nickname,
		Role:     "user",
		Status:   "active",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceImpl) Login(req LoginRequest) (string, error) {
	//查找用户
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return "", errors.New("用户名或密码错误")
	}

	//检查用户状态
	if user.Status != "active" {
		return "", errors.New("用户已被封禁")
	}

	//生成token
	token, err := auth.GenerateToken(user.ID, user.Username, user.Role, 24*time.Hour)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return token, nil
}

func (s *UserServiceImpl) GetProfile(userID uint, req UpdateProfileRequest) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}

	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}

	return s.userRepo.Update(user)
}
