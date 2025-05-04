package dto

import (
	model "LmsSystem/models"
	"time"
)

// LessonRequest represents the request for creating or updating a lesson
type LessonRequest struct {
	Name        string `json:"name" binding:"required,max=255"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Order       int    `json:"order" binding:"required"`
	ChapterID   uint   `json:"chapter_id" binding:"required"`
}

// LessonResponse represents the response for lesson data
type LessonResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Order       int       `json:"order"`
	ChapterID   uint      `json:"chapter_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FromModel converts a model.Lesson to LessonResponse
func (lr *LessonResponse) FromModel(lesson *model.Lesson) {
	lr.ID = lesson.ID
	lr.Name = lesson.Name
	lr.Description = lesson.Description
	lr.Content = lesson.Content
	lr.Order = lesson.Order
	lr.ChapterID = lesson.ChapterID
	lr.CreatedAt = lesson.CreatedAt
	lr.UpdatedAt = lesson.UpdatedAt
}

// ToModel converts a LessonRequest to model.Lesson
func (lr *LessonRequest) ToModel() *model.Lesson {
	return &model.Lesson{
		Name:        lr.Name,
		Description: lr.Description,
		Content:     lr.Content,
		Order:       lr.Order,
		ChapterID:   lr.ChapterID,
	}
}
