package model

import (
	"time"

	"gorm.io/gorm"
)

// Chapter represents a chapter entity within a course
type Chapter struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Order       int            `gorm:"not null" json:"order"`
	CourseID    uint           `gorm:"not null" json:"course_id"`
	Course      Course         `gorm:"foreignKey:CourseID" json:"-"`
	Lessons     []Lesson       `gorm:"foreignKey:ChapterID;constraint:OnDelete:CASCADE" json:"lessons,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
