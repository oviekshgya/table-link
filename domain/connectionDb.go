package domain

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"table-link/src/model/role"
	"table-link/src/model/users"
	"time"
)

type DatabaseConfig struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         int
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

var DB *gorm.DB

func (conf DatabaseConfig) ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", conf.Host, conf.Username, conf.Password, conf.Dbname, conf.Port)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal menghubungkan ke database!")
	}

	db.AutoMigrate(&users.Users{}, role.Role{}, role.RoleRight{})
	DB = db

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan instance *sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)
}
