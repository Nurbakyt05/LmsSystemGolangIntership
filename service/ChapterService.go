package service

import (
	"LmsSystem/models"
	"LmsSystem/repository"
)

type ChapterService interface {
	GetAllChapters() ([]models.Chapter, error)
	GetChapterByID(id uint) (*models.Chapter, error)
	GetChaptersByCourse(courseID uint) ([]models.Chapter, error)
	CreateChapter(chapter *models.Chapter) error
	UpdateChapter(chapter *models.Chapter) error
	DeleteChapter(id uint) error
}

type chapterService struct {
	repo repository.ChapterRepository
}

func NewChapterService(repo repository.ChapterRepository) ChapterService {
	return &chapterService{repo: repo}
}

func (s *chapterService) GetAllChapters() ([]models.Chapter, error) {
	return s.repo.GetAll()
}

func (s *chapterService) GetChapterByID(id uint) (*models.Chapter, error) {
	return s.repo.GetByID(id)
}

func (s *chapterService) GetChaptersByCourse(courseID uint) ([]models.Chapter, error) {
	return s.repo.GetByCourseID(courseID)
}

func (s *chapterService) CreateChapter(chapter *models.Chapter) error {
	return s.repo.Create(chapter)
}

func (s *chapterService) UpdateChapter(chapter *models.Chapter) error {
	return s.repo.Update(chapter)
}

func (s *chapterService) DeleteChapter(id uint) error {
	return s.repo.Delete(id)
}
