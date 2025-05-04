package service

import (
	"LmsSystem/dto"
	"LmsSystem/repository"
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// LessonService defines the interface for lesson business logic
type LessonService interface {
	CreateLesson(ctx context.Context, lessonReq *dto.LessonRequest) (*dto.LessonResponse, error)
	UpdateLesson(ctx context.Context, id uint, lessonReq *dto.LessonRequest) (*dto.LessonResponse, error)
	DeleteLesson(ctx context.Context, id uint) error
	GetLessonByID(ctx context.Context, id uint) (*dto.LessonResponse, error)
	GetLessonsByChapterID(ctx context.Context, chapterID uint) ([]dto.LessonResponse, error)
}

// lessonService implements LessonService interface
type lessonService struct {
	repo        repository.LessonRepository
	chapterRepo repository.ChapterRepository
	logger      *logrus.Logger
}

// NewLessonService creates a new instance of LessonService
func NewLessonService(lessonRepo repository.LessonRepository, chapterRepo repository.ChapterRepository, logger *logrus.Logger) LessonService {
	return &lessonService{
		repo:        lessonRepo,
		chapterRepo: chapterRepo,
		logger:      logger,
	}
}

// CreateLesson creates a new lesson
func (s *lessonService) CreateLesson(ctx context.Context, lessonReq *dto.LessonRequest) (*dto.LessonResponse, error) {
	// Verify chapter exists
	_, err := s.chapterRepo.FindByID(ctx, lessonReq.ChapterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("chapter_id", lessonReq.ChapterID).Info("Chapter not found for lesson creation")
			return nil, fmt.Errorf("chapter with ID %d not found", lessonReq.ChapterID)
		}
		s.logger.WithError(err).WithField("chapter_id", lessonReq.ChapterID).Error("Error finding chapter for lesson creation")
		return nil, err
	}

	lesson := lessonReq.ToModel()
	err = s.repo.Create(ctx, lesson)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create lesson")
		return nil, err
	}

	s.logger.WithFields(logrus.Fields{
		"lesson_id":   lesson.ID,
		"lesson_name": lesson.Name,
		"chapter_id":  lesson.ChapterID,
	}).Info("Lesson created successfully")

	var resp dto.LessonResponse
	resp.FromModel(lesson)
	return &resp, nil
}

// UpdateLesson updates an existing lesson
func (s *lessonService) UpdateLesson(ctx context.Context, id uint, lessonReq *dto.LessonRequest) (*dto.LessonResponse, error) {
	existingLesson, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("lesson_id", id).Info("Lesson not found for update")
			return nil, fmt.Errorf("lesson with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("lesson_id", id).Error("Error finding lesson for update")
		return nil, err
	}

	// If chapter ID is changing, verify new chapter exists
	if existingLesson.ChapterID != lessonReq.ChapterID {
		_, err := s.chapterRepo.FindByID(ctx, lessonReq.ChapterID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				s.logger.WithField("chapter_id", lessonReq.ChapterID).Info("Chapter not found for lesson update")
				return nil, fmt.Errorf("chapter with ID %d not found", lessonReq.ChapterID)
			}
			s.logger.WithError(err).WithField("chapter_id", lessonReq.ChapterID).Error("Error finding chapter for lesson update")
			return nil, err
		}
	}

	existingLesson.Name = lessonReq.Name
	existingLesson.Description = lessonReq.Description
	existingLesson.Content = lessonReq.Content
	existingLesson.Order = lessonReq.Order
	existingLesson.ChapterID = lessonReq.ChapterID

	err = s.repo.Update(ctx, existingLesson)
	if err != nil {
		s.logger.WithError(err).WithField("lesson_id", id).Error("Failed to update lesson")
		return nil, err
	}

	s.logger.WithField("lesson_id", id).Info("Lesson updated successfully")

	var resp dto.LessonResponse
	resp.FromModel(existingLesson)
	return &resp, nil
}

// DeleteLesson deletes a lesson by ID
func (s *lessonService) DeleteLesson(ctx context.Context, id uint) error {
	// Check if lesson exists
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("lesson_id", id).Info("Lesson not found for deletion")
			return fmt.Errorf("lesson with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("lesson_id", id).Error("Error finding lesson for deletion")
		return err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.WithError(err).WithField("lesson_id", id).Error("Failed to delete lesson")
		return err
	}

	s.logger.WithField("lesson_id", id).Info("Lesson deleted successfully")
	return nil
}

// GetLessonByID retrieves a lesson by ID
func (s *lessonService) GetLessonByID(ctx context.Context, id uint) (*dto.LessonResponse, error) {
	lesson, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("lesson_id", id).Info("Lesson not found")
			return nil, fmt.Errorf("lesson with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("lesson_id", id).Error("Error finding lesson")
		return nil, err
	}

	var resp dto.LessonResponse
	resp.FromModel(lesson)
	return &resp, nil
}

// GetLessonsByChapterID retrieves all lessons for a specific chapter
func (s *lessonService) GetLessonsByChapterID(ctx context.Context, chapterID uint) ([]dto.LessonResponse, error) {
	// Verify chapter exists
	_, err := s.chapterRepo.FindByID(ctx, chapterID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("chapter_id", chapterID).Info("Chapter not found for lesson retrieval")
			return nil, fmt.Errorf("chapter with ID %d not found", chapterID)
		}
		s.logger.WithError(err).WithField("chapter_id", chapterID).Error("Error finding chapter for lesson retrieval")
		return nil, err
	}

	lessons, err := s.repo.FindByChapterID(ctx, chapterID)
	if err != nil {
		s.logger.WithError(err).WithField("chapter_id", chapterID).Error("Failed to retrieve lessons for chapter")
		return nil, err
	}

	responses := make([]dto.LessonResponse, len(lessons))
	for i, lesson := range lessons {
		responses[i].FromModel(&lesson)
	}
	return responses, nil
}
