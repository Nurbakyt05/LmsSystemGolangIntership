package lesson

import (
	"LmsSystem/models"
	"gorm.io/gorm"
)

type LessonRepository interface {
	GetAll() ([]models.Lesson, error)
	GetByID(id uint) (*models.Lesson, error)
	Create(lesson *models.Lesson) error
	Update(lesson *models.Lesson) error
	Delete(id uint) error
}

type lessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) LessonRepository {
	return &lessonRepository{db}
}

func (r *lessonRepository) GetAll() ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := r.db.Preload("Chapter").Find(&lessons).Error
	return lessons, err
}

func (r *lessonRepository) GetByID(id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.Preload("Chapter").First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *lessonRepository) Create(lesson *models.Lesson) error {
	return r.db.Create(lesson).Error
}

func (r *lessonRepository) Update(lesson *models.Lesson) error {
	return r.db.Save(lesson).Error
}

func (r *lessonRepository) Delete(id uint) error {
	return r.db.Delete(&models.Lesson{}, id).Error
}
