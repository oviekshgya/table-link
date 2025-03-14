package role

import (
	"gorm.io/gorm"
	"log"
)

type RoleModel struct {
	DB *gorm.DB
}

const ROLES = "roles"
const ROLERIGHT = "role_rights"

type Role struct {
	ID         uint        `gorm:"primaryKey"`
	Name       string      `gorm:"unique;not null"`
	RoleRights []RoleRight `gorm:"foreignKey:RoleID" json:"role_rights"`
}

func (Role) TableName() string {
	return ROLES
}

type RoleRight struct {
	ID      uint   `gorm:"primaryKey"`
	RoleID  uint   `gorm:"not null"`
	Section string `gorm:"not null"`
	Route   string `gorm:"not null"`
	RCreate bool   `gorm:"default:false"`
	RRead   bool   `gorm:"default:false"`
	RUpdate bool   `gorm:"default:false"`
}

func (RoleRight) TableName() string {
	return ROLERIGHT
}

func (repo *RoleModel) GetRoleWithRights() ([]*Role, error) {
	var role []*Role
	err := repo.DB.Preload("RoleRights").Find(&role).Error
	if err != nil {
		log.Println("Gagal mengambil role dengan hak akses:", err)
		return nil, err
	}
	return role, nil
}
