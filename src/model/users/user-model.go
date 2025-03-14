package users

import (
	"gorm.io/gorm"
	"table-link/src/model/role"
)

const USERS = "users"

type UserModel struct {
	DB *gorm.DB
}

type Users struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string `gorm:"column:password"`
	RoleID   uint
	Role     role.Role
}

func (Users) TableName() string {
	return USERS
}

func (s *UserModel) Create(data Users) error {
	return s.DB.Create(&data).Error
}

func (s *UserModel) GetByUsername(username string) (*Users, error) {
	var data Users
	err := s.DB.Where("username = ?", username).First(&data).Error
	return &data, err
}

func (s *UserModel) GetAll() ([]Users, error) {
	var users []Users
	err := s.DB.Preload("Role").Find(&users).Error
	return users, err
}
