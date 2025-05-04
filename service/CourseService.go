package service

import (
	"LmsSystem/models"
	"LmsSystem/repository"
)

type CourseService interface {
	GetAll() ([]models.Course, error)
	GetByID(id uint) (*models.Course, error)
	GetFullCourse(id uint) (*models.Course, error)
	Create(course *models.Course) error
	Update(course *models.Course) error
	Delete(id uint) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) GetAll() ([]models.Course, error) {
	return s.repo.GetAll()
}

func (s *courseService) GetByID(id uint) (*models.Course, error) {
	return s.repo.GetByID(id)
}

func (s *courseService) GetFullCourse(id uint) (*models.Course, error) {
	return s.repo.GetWithChapters(id)
}

func (s *courseService) Create(course *models.Course) error {
	return s.repo.Create(course)
}

func (s *courseService) Update(course *models.Course) error {
	return s.repo.Update(course)
}

func (s *courseService) Delete(id uint) error {
	return s.repo.Delete(id)
}
