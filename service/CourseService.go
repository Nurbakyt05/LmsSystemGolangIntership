package service

import (
	"LmsSystem/dto"
	course2 "LmsSystem/mapper"
	course3 "LmsSystem/repository"
)

type CourseService interface {
	GetAll() ([]dto.CourseDTO, error)
	GetByID(id uint) (*dto.CourseDTO, error)
	Create(course dto.CourseDTO) error
	Update(course dto.CourseDTO) error
	Delete(id uint) error
}

type courseService struct {
	repo course3.CourseRepository
}

func NewCourseService(repo course3.CourseRepository) CourseService {
	return &courseService{repo}
}

func (s *courseService) GetAll() ([]dto.CourseDTO, error) {
	courses, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.CourseDTO
	for _, course := range courses {
		result = append(result, course2.ToCourseDTO(course))
	}
	return result, nil
}

func (s *courseService) GetByID(id uint) (*dto.CourseDTO, error) {
	course, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	dto := course2.ToCourseDTO(*course)
	return &dto, nil
}

func (s *courseService) Create(course dto.CourseDTO) error {
	model := course2.ToCourseModel(course)
	return s.repo.Create(&model)
}

func (s *courseService) Update(course dto.CourseDTO) error {
	model := course2.ToCourseModel(course)
	return s.repo.Update(&model)
}

func (s *courseService) Delete(id uint) error {
	return s.repo.Delete(id)
}
