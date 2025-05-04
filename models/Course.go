// models/course.go
package models

import "time"

type Course struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Chapters    []Chapter `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE" json:"chapters,omitempty"`
}
