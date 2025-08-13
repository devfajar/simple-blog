package database

import (
	"fmt"
	"log"

	"github.com/devfajar/blog-app/internal/config"
	"github.com/devfajar/blog-app/internal/domain"
	"github.com/devfajar/blog-app/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{}, &domain.Category{}, &domain.Post{})
}

func SeedAdmin(db *gorm.DB, cfg config.Config) error {
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		return nil
	}
	hash, _ := util.HashPassword(cfg.AdminPass)
	admin := domain.User{
		Name:         cfg.AdminName,
		Email:        cfg.AdminEmail,
		PasswordHash: hash,
		Role:         "admin",
	}
	return db.Create(&admin).Error
}
