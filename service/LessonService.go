package service

import (
	"LmsSystem/models"
	"LmsSystem/repository"
)

type LessonService interface {
	GetAllLessons() ([]models.Lesson, error)
	GetLessonByID(id uint) (*models.Lesson, error)
	GetLessonsByChapter(chapterID uint) ([]models.Lesson, error)
	CreateLesson(lesson *models.Lesson) error
	UpdateLesson(lesson *models.Lesson) error
	DeleteLesson(id uint) error
}

type lessonService struct {
	repo repository.LessonRepository
}

func NewLessonService(repo repository.LessonRepository) LessonService {
	return &lessonService{repo: repo}
}

func (s *lessonService) GetAllLessons() ([]models.Lesson, error) {
	return s.repo.GetAll()
}

func (s *lessonService) GetLessonByID(id uint) (*models.Lesson, error) {
	return s.repo.GetByID(id)
}

func (s *lessonService) GetLessonsByChapter(chapterID uint) ([]models.Lesson, error) {
	return s.repo.GetByChapterID(chapterID)
}

func (s *lessonService) CreateLesson(lesson *models.Lesson) error {
	return s.repo.Create(lesson)
}

func (s *lessonService) UpdateLesson(lesson *models.Lesson) error {
	return s.repo.Update(lesson)
}

func (s *lessonService) DeleteLesson(id uint) error {
	return s.repo.Delete(id)
}
