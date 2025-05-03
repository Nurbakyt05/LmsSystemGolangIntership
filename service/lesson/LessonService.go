package lesson

import (
	dto "LmsSystem/dto/lesson"
	mapper "LmsSystem/mapper/lesson"
	repo "LmsSystem/repository/lesson"
)

type LessonService interface {
	GetAll() ([]dto.LessonDTO, error)
	GetByID(id uint) (*dto.LessonDTO, error)
	Create(dto.LessonDTO) error
	Update(dto.LessonDTO) error
	Delete(id uint) error
}

type lessonService struct {
	repo repo.LessonRepository
}

func NewLessonService(repo repo.LessonRepository) LessonService {
	return &lessonService{repo}
}

func (s *lessonService) GetAll() ([]dto.LessonDTO, error) {
	lessons, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.LessonDTO
	for _, l := range lessons {
		result = append(result, mapper.ToLessonDTO(l))
	}
	return result, nil
}

func (s *lessonService) GetByID(id uint) (*dto.LessonDTO, error) {
	l, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	dtoObj := mapper.ToLessonDTO(*l)
	return &dtoObj, nil
}

func (s *lessonService) Create(dto dto.LessonDTO) error {
	model := mapper.ToLessonModel(dto)
	return s.repo.Create(&model)
}

func (s *lessonService) Update(dto dto.LessonDTO) error {
	model := mapper.ToLessonModel(dto)
	return s.repo.Update(&model)
}

func (s *lessonService) Delete(id uint) error {
	return s.repo.Delete(id)
}
