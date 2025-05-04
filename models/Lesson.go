package model

import (
	"time"

	"gorm.io/gorm"
)

// Lesson represents a lesson entity within a chapter
type Lesson struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Content     string         `gorm:"type:text" json:"content"`
	Order       int            `gorm:"not null" json:"order"`
	ChapterID   uint           `gorm:"not null" json:"chapter_id"`
	Chapter     Chapter        `gorm:"foreignKey:ChapterID" json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
