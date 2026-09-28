package repository

import (
	"blog-system/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	ExistsByUsername(username string) bool
	ExistsByEmail(email string) bool
	FindByUsername(username string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	Update(user *model.User) error
}

type UserServiceImpl struct {
	userRepo *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserServiceImpl{userRepo: db}
}

func (r *UserServiceImpl) Create(user *model.User) error {
	return r.userRepo.Create(user).Error
}

func (r *UserServiceImpl) ExistsByUsername(username string) bool {
	var count int64
	r.userRepo.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (r *UserServiceImpl) ExistsByEmail(email string) bool {
	var count int64
	r.userRepo.Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *UserServiceImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.userRepo.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserServiceImpl) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.userRepo.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserServiceImpl) Update(user *model.User) error {
	return r.userRepo.Save(user).Error
}
