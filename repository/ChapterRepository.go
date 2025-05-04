package repository

import (
	"LmsSystem/models"
	"gorm.io/gorm"
)

type ChapterRepository interface {
	GetAll() ([]models.Chapter, error)
	GetByID(id uint) (*models.Chapter, error)
	GetByCourseID(courseID uint) ([]models.Chapter, error)
	Create(chapter *models.Chapter) error
	Update(chapter *models.Chapter) error
	Delete(id uint) error
}

type chapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{db: db}
}

func (r *chapterRepository) GetAll() ([]models.Chapter, error) {
	var chapters []models.Chapter
	err := r.db.Preload("Lessons").Find(&chapters).Error
	return chapters, err
}

func (r *chapterRepository) GetByID(id uint) (*models.Chapter, error) {
	var chapter models.Chapter
	err := r.db.Preload("Lessons").First(&chapter, id).Error
	if err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (r *chapterRepository) GetByCourseID(courseID uint) ([]models.Chapter, error) {
	var chapters []models.Chapter
	err := r.db.Preload("Lessons").Where("course_id = ?", courseID).Find(&chapters).Error
	return chapters, err
}

func (r *chapterRepository) Create(chapter *models.Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *chapterRepository) Update(chapter *models.Chapter) error {
	return r.db.Save(chapter).Error
}

func (r *chapterRepository) Delete(id uint) error {
	return r.db.Delete(&models.Chapter{}, id).Error
}
