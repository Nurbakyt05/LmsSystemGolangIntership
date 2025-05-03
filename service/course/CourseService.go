package service

import (
	"LmsSystem/dto"
	"LmsSystem/mapper"
	"LmsSystem/repository"
)

type CourseService interface {
	GetAll() ([]dto.CourseDTO, error)
	GetByID(id uint) (*dto.CourseDTO, error)
	Create(course dto.CourseDTO) error
	Update(course dto.CourseDTO) error
	Delete(id uint) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo}
}

func (s *courseService) GetAll() ([]dto.CourseDTO, error) {
	courses, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.CourseDTO
	for _, course := range courses {
		result = append(result, mapper.ToCourseDTO(course))
	}
	return result, nil
}

func (s *courseService) GetByID(id uint) (*dto.CourseDTO, error) {
	course, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	dto := mapper.ToCourseDTO(*course)
	return &dto, nil
}

func (s *courseService) Create(course dto.CourseDTO) error {
	model := mapper.ToCourseModel(course)
	return s.repo.Create(&model)
}

func (s *courseService) Update(course dto.CourseDTO) error {
	model := mapper.ToCourseModel(course)
	return s.repo.Update(&model)
}

func (s *courseService) Delete(id uint) error {
	return s.repo.Delete(id)
}
