package model

type UserProfile struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"uniqueIndex"`
	User   User `gorm:"foreignKey:UserID"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
