package service

import (
	"LmsSystem/dto"
	"LmsSystem/mapper"
	"LmsSystem/repository"
)

type ChapterService interface {
	GetAll() ([]dto.ChapterDTO, error)
	GetByID(id uint) (*dto.ChapterDTO, error)
	Create(dto dto.ChapterDTO) error
	Update(dto dto.ChapterDTO) error
	Delete(id uint) error
}

type chapterService struct {
	repo repository.ChapterRepository
}

func NewChapterService(repo repository.ChapterRepository) ChapterService {
	return &chapterService{repo}
}

func (s *chapterService) GetAll() ([]dto.ChapterDTO, error) {
	chapters, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []dto.ChapterDTO
	for _, c := range chapters {
		result = append(result, mapper.ToChapterDTO(c))
	}
	return result, nil
}

func (s *chapterService) GetByID(id uint) (*dto.ChapterDTO, error) {
	ch, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	dto := mapper.ToChapterDTO(*ch)
	return &dto, nil
}

func (s *chapterService) Create(dto dto.ChapterDTO) error {
	model := mapper.ToChapterModel(dto)
	return s.repo.Create(&model)
}

func (s *chapterService) Update(dto dto.ChapterDTO) error {
	model := mapper.ToChapterModel(dto)
	return s.repo.Update(&model)
}

func (s *chapterService) Delete(id uint) error {
	return s.repo.Delete(id)
}
