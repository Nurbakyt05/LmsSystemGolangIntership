package dto

import (
	model "LmsSystem/models"
	"time"
)

// CourseRequest represents the request for creating or updating a course
type CourseRequest struct {
	Name        string `json:"name" binding:"required,max=255"`
	Description string `json:"description"`
}

// CourseResponse represents the response for course data
type CourseResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CourseDetailResponse represents the response for detailed course data including chapters
type CourseDetailResponse struct {
	CourseResponse
	Chapters []ChapterResponse `json:"chapters,omitempty"`
}

// FromModel converts a model.Course to CourseResponse
func (cr *CourseResponse) FromModel(course *model.Course) {
	cr.ID = course.ID
	cr.Name = course.Name
	cr.Description = course.Description
	cr.CreatedAt = course.CreatedAt
	cr.UpdatedAt = course.UpdatedAt
}

// FromModel converts a model.Course to CourseDetailResponse including chapters
func (cdr *CourseDetailResponse) FromModel(course *model.Course) {
	cdr.ID = course.ID
	cdr.Name = course.Name
	cdr.Description = course.Description
	cdr.CreatedAt = course.CreatedAt
	cdr.UpdatedAt = course.UpdatedAt

	if len(course.Chapters) > 0 {
		cdr.Chapters = make([]ChapterResponse, len(course.Chapters))
		for i, chapter := range course.Chapters {
			cdr.Chapters[i].FromModel(&chapter)
		}
	}
}

// ToModel converts a CourseRequest to model.Course
func (cr *CourseRequest) ToModel() *model.Course {
	return &model.Course{
		Name:        cr.Name,
		Description: cr.Description,
	}
}
