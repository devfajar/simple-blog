package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `json:"name"`
	Email        string         `gorm:"uniqueIndex;size:191" json:"email"`
	PasswordHash string         `json:"-"`
	Role         string         `json:"role"` // "admin"
}
