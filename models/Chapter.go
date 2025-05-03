package models

import "time"

type Chapter struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Order       int
	CourseID    uint
	Course      Course `gorm:"foreignKey:CourseID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
