package models

import "time"

type Course struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
