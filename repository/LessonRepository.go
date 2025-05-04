package repository

import (
	model "LmsSystem/models"
	"context"
	"gorm.io/gorm"
)

// LessonRepository defines the interface for lesson data access
type LessonRepository interface {
	Create(ctx context.Context, lesson *model.Lesson) error
	Update(ctx context.Context, lesson *model.Lesson) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*model.Lesson, error)
	FindByChapterID(ctx context.Context, chapterID uint) ([]model.Lesson, error)
}

// lessonRepository implements LessonRepository interface
type lessonRepository struct {
	db *gorm.DB
}

// NewLessonRepository creates a new instance of LessonRepository
func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &lessonRepository{
		db: db,
	}
}

// Create adds a new lesson to the database
func (r *lessonRepository) Create(ctx context.Context, lesson *model.Lesson) error {
	return r.db.WithContext(ctx).Create(lesson).Error
}

// Update updates an existing lesson in the database
func (r *lessonRepository) Update(ctx context.Context, lesson *model.Lesson) error {
	return r.db.WithContext(ctx).Save(lesson).Error
}

// Delete removes a lesson from the database by ID
func (r *lessonRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Lesson{}, id).Error
}

// FindByID retrieves a lesson by its ID
func (r *lessonRepository) FindByID(ctx context.Context, id uint) (*model.Lesson, error) {
	var lesson model.Lesson
	err := r.db.WithContext(ctx).First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

// FindByChapterID retrieves all lessons for a specific chapter ID
func (r *lessonRepository) FindByChapterID(ctx context.Context, chapterID uint) ([]model.Lesson, error) {
	var lessons []model.Lesson
	err := r.db.WithContext(ctx).Where("chapter_id = ?", chapterID).Order("\"order\" ASC").Find(&lessons).Error
	if err != nil {
		return nil, err
	}
	return lessons, nil
}
