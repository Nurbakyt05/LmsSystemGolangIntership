package repository

import (
	model "LmsSystem/models"
	"context"

	"gorm.io/gorm"
)

// CourseRepository defines the interface for course data access
type CourseRepository interface {
	Create(ctx context.Context, course *model.Course) error
	Update(ctx context.Context, course *model.Course) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*model.Course, error)
	FindAll(ctx context.Context) ([]model.Course, error)
	FindByIDWithChapters(ctx context.Context, id uint) (*model.Course, error)
}

// courseRepository implements CourseRepository interface
type courseRepository struct {
	db *gorm.DB
}

// NewCourseRepository creates a new instance of CourseRepository
func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}

// Create adds a new course to the database
func (r *courseRepository) Create(ctx context.Context, course *model.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

// Update updates an existing course in the database
func (r *courseRepository) Update(ctx context.Context, course *model.Course) error {
	return r.db.WithContext(ctx).Save(course).Error
}

// Delete removes a course from the database by ID
func (r *courseRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Course{}, id).Error
}

// FindByID retrieves a course by its ID
func (r *courseRepository) FindByID(ctx context.Context, id uint) (*model.Course, error) {
	var course model.Course
	err := r.db.WithContext(ctx).First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

// FindAll retrieves all courses
func (r *courseRepository) FindAll(ctx context.Context) ([]model.Course, error) {
	var courses []model.Course
	err := r.db.WithContext(ctx).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// FindByIDWithChapters retrieves a course with its chapters by ID
func (r *courseRepository) FindByIDWithChapters(ctx context.Context, id uint) (*model.Course, error) {
	var course model.Course
	err := r.db.WithContext(ctx).Preload("Chapters", func(db *gorm.DB) *gorm.DB {
		return db.Order("chapters.order ASC")
	}).First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}
