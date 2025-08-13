package domain

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"index;size:200" json:"title"`
	Slug        string         `gorm:"uniqueIndex;size:191" json:"slug"`
	Content     string         `gorm:"type:longtext" json:"content"`
	Status      string         `gorm:"index;size:20" json:"status"` // draft/published
	PublishedAt *time.Time     `json:"published_at,omitempty"`
	AuthorID    uint           `json:"author_id"`
	Author      User           `gorm:"foreignKey:AuthorID" json:"-"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"-"`
}
