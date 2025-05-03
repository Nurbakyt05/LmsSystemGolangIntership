package models

import "time"

type Chapter struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Order       int       `json:"order"`
	CourseID    uint      `json:"course_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
