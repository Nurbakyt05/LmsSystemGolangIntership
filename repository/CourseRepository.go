package repository

import (
	"LmsSystem/models"
	"gorm.io/gorm"
)

type CourseRepository interface {
	GetAll() ([]models.Course, error)
	GetByID(id uint) (*models.Course, error)
	GetWithChapters(id uint) (*models.Course, error)
	Create(course *models.Course) error
	Update(course *models.Course) error
	Delete(id uint) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) GetAll() ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Find(&courses).Error
	return courses, err
}

func (r *courseRepository) GetByID(id uint) (*models.Course, error) {
	var course models.Course
	err := r.db.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) GetWithChapters(id uint) (*models.Course, error) {
	var course models.Course
	err := r.db.Preload("Chapters.Lessons").First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) Update(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepository) Delete(id uint) error {
	return r.db.Select("Chapters").Delete(&models.Course{}, id).Error
}
