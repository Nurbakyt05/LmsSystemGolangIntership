package dto

import (
	model "LmsSystem/models"
	"time"
)

// ChapterRequest represents the request for creating or updating a chapter
type ChapterRequest struct {
	Name        string `json:"name" binding:"required,max=255"`
	Description string `json:"description"`
	Order       int    `json:"order" binding:"required"`
	CourseID    uint   `json:"course_id" binding:"required"`
}

// ChapterResponse represents the response for chapter data
type ChapterResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	CourseID    uint      `json:"course_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ChapterDetailResponse represents the response for detailed chapter data including lessons
type ChapterDetailResponse struct {
	ChapterResponse
	Lessons []LessonResponse `json:"lessons,omitempty"`
}

// FromModel converts a model.Chapter to ChapterResponse
func (cr *ChapterResponse) FromModel(chapter *model.Chapter) {
	cr.ID = chapter.ID
	cr.Name = chapter.Name
	cr.Description = chapter.Description
	cr.Order = chapter.Order
	cr.CourseID = chapter.CourseID
	cr.CreatedAt = chapter.CreatedAt
	cr.UpdatedAt = chapter.UpdatedAt
}

// FromModel converts a model.Chapter to ChapterDetailResponse including lessons
func (cdr *ChapterDetailResponse) FromModel(chapter *model.Chapter) {
	cdr.ID = chapter.ID
	cdr.Name = chapter.Name
	cdr.Description = chapter.Description
	cdr.Order = chapter.Order
	cdr.CourseID = chapter.CourseID
	cdr.CreatedAt = chapter.CreatedAt
	cdr.UpdatedAt = chapter.UpdatedAt

	if len(chapter.Lessons) > 0 {
		cdr.Lessons = make([]LessonResponse, len(chapter.Lessons))
		for i, lesson := range chapter.Lessons {
			cdr.Lessons[i].FromModel(&lesson)
		}
	}
}

// ToModel converts a ChapterRequest to model.Chapter
func (cr *ChapterRequest) ToModel() *model.Chapter {
	return &model.Chapter{
		Name:        cr.Name,
		Description: cr.Description,
		Order:       cr.Order,
		CourseID:    cr.CourseID,
	}
}
