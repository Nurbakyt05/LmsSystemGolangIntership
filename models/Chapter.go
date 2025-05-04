// models/chapter.go
package models

import "time"

type Chapter struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Order       int       `json:"order"`
	CourseID    uint      `json:"courseId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Lessons     []Lesson  `gorm:"foreignKey:ChapterID;constraint:OnDelete:CASCADE" json:"lessons,omitempty"`
}
