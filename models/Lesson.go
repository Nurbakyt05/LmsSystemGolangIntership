package models

import "time"

type Lesson struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Content     string `gorm:"type:text"`
	Order       int
	ChapterID   uint
	Chapter     Chapter `gorm:"foreignKey:ChapterID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
