package users

import (
	"gorm.io/gorm"
	"table-link/src/model/role"
)

const USERS = "users"

type UserModel struct {
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

func (s *UserModel) Create(db *gorm.DB, data Users) error {
	return db.Create(&data).Error
}

func (s *UserModel) GetByUsername(db *gorm.DB, username string) (*Users, error) {
	var data Users
	err := db.Where("username = ?", username).First(&data).Error
	return &data, err
}

func (s *UserModel) GetAll(db *gorm.DB) ([]Users, error) {
	var users []Users
	err := db.Preload("Role").Find(&users).Error
	return users, err
}

func (r *UserModel) DeleteUserByID(db *gorm.DB, userID uint) error {
	result := db.Delete(&Users{}, userID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *UserModel) Update(db *gorm.DB, name string) error {
	result := db.Table(USERS).Where("name = ?", name).Updates(map[string]interface{}{
		"name": name,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
