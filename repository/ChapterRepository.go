package repository

import (
	model "LmsSystem/models"
	"context"
	"gorm.io/gorm"
)

// ChapterRepository defines the interface for chapter data access
type ChapterRepository interface {
	Create(ctx context.Context, chapter *model.Chapter) error
	Update(ctx context.Context, chapter *model.Chapter) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*model.Chapter, error)
	FindByCourseID(ctx context.Context, courseID uint) ([]model.Chapter, error)
	FindByIDWithLessons(ctx context.Context, id uint) (*model.Chapter, error)
}

// chapterRepository implements ChapterRepository interface
type chapterRepository struct {
	db *gorm.DB
}

// NewChapterRepository creates a new instance of ChapterRepository
func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{
		db: db,
	}
}

// Create adds a new chapter to the database
func (r *chapterRepository) Create(ctx context.Context, chapter *model.Chapter) error {
	return r.db.WithContext(ctx).Create(chapter).Error
}

// Update updates an existing chapter in the database
func (r *chapterRepository) Update(ctx context.Context, chapter *model.Chapter) error {
	return r.db.WithContext(ctx).Save(chapter).Error
}

// Delete removes a chapter from the database by ID
func (r *chapterRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Chapter{}, id).Error
}

// FindByID retrieves a chapter by its ID
func (r *chapterRepository) FindByID(ctx context.Context, id uint) (*model.Chapter, error) {
	var chapter model.Chapter
	err := r.db.WithContext(ctx).First(&chapter, id).Error
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}

// FindByCourseID retrieves all chapters for a specific course ID
func (r *chapterRepository) FindByCourseID(ctx context.Context, courseID uint) ([]model.Chapter, error) {
	var chapters []model.Chapter
	err := r.db.WithContext(ctx).Where("course_id = ?", courseID).Order("\"order\" ASC").Find(&chapters).Error
	if err != nil {
		return nil, err
	}
	return chapters, nil
}

// FindByIDWithLessons retrieves a chapter with its lessons by ID
func (r *chapterRepository) FindByIDWithLessons(ctx context.Context, id uint) (*model.Chapter, error) {
	var chapter model.Chapter
	err := r.db.WithContext(ctx).
		Preload("Lessons", func(db *gorm.DB) *gorm.DB {
			return db.Order("lessons.order ASC")
		}).
		First(&chapter, id).
		Error
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}
