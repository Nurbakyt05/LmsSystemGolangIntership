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

// ChapterService defines the interface for chapter business logic
type ChapterService interface {
	CreateChapter(ctx context.Context, chapterReq *dto.ChapterRequest) (*dto.ChapterResponse, error)
	UpdateChapter(ctx context.Context, id uint, chapterReq *dto.ChapterRequest) (*dto.ChapterResponse, error)
	DeleteChapter(ctx context.Context, id uint) error
	GetChapterByID(ctx context.Context, id uint) (*dto.ChapterResponse, error)
	GetChaptersByCourseID(ctx context.Context, courseID uint) ([]dto.ChapterResponse, error)
	GetChapterWithLessons(ctx context.Context, id uint) (*dto.ChapterDetailResponse, error)
}

type chapterService struct {
	repo       repository.ChapterRepository
	courseRepo repository.CourseRepository
	logger     *logrus.Logger
}

// NewChapterService creates a new instance of ChapterService
func NewChapterService(chapterRepo repository.ChapterRepository, courseRepo repository.CourseRepository, logger *logrus.Logger) ChapterService {
	return &chapterService{
		repo:       chapterRepo,
		courseRepo: courseRepo,
		logger:     logger,
	}
}

// CreateChapter creates a new chapter
func (s *chapterService) CreateChapter(ctx context.Context, chapterReq *dto.ChapterRequest) (*dto.ChapterResponse, error) {
	// Verify course exists
	_, err := s.courseRepo.FindByID(ctx, chapterReq.CourseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("course_id", chapterReq.CourseID).Info("Course not found for chapter creation")
			return nil, fmt.Errorf("course with ID %d not found", chapterReq.CourseID)
		}
		s.logger.WithError(err).WithField("course_id", chapterReq.CourseID).Error("Error finding course for chapter creation")
		return nil, err
	}

	chapter := chapterReq.ToModel()
	err = s.repo.Create(ctx, chapter)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create chapter")
		return nil, err
	}

	s.logger.WithFields(logrus.Fields{
		"chapter_id":   chapter.ID,
		"chapter_name": chapter.Name,
		"course_id":    chapter.CourseID,
	}).Info("Chapter created successfully")

	var resp dto.ChapterResponse
	resp.FromModel(chapter)
	return &resp, nil
}

// UpdateChapter updates an existing chapter
func (s *chapterService) UpdateChapter(ctx context.Context, id uint, chapterReq *dto.ChapterRequest) (*dto.ChapterResponse, error) {
	existingChapter, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("chapter_id", id).Info("Chapter not found for update")
			return nil, fmt.Errorf("chapter with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("chapter_id", id).Error("Error finding chapter for update")
		return nil, err
	}

	if existingChapter.CourseID != chapterReq.CourseID {
		_, err := s.courseRepo.FindByID(ctx, chapterReq.CourseID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				s.logger.WithField("course_id", chapterReq.CourseID).Info("Course not found for chapter update")
				return nil, fmt.Errorf("course with ID %d not found", chapterReq.CourseID)
			}
			s.logger.WithError(err).WithField("course_id", chapterReq.CourseID).Error("Error finding course for chapter update")
			return nil, err
		}
	}

	existingChapter.Name = chapterReq.Name
	existingChapter.Description = chapterReq.Description
	existingChapter.Order = chapterReq.Order
	existingChapter.CourseID = chapterReq.CourseID

	err = s.repo.Update(ctx, existingChapter)
	if err != nil {
		s.logger.WithError(err).WithField("chapter_id", id).Error("Failed to update chapter")
		return nil, err
	}

	s.logger.WithField("chapter_id", id).Info("Chapter updated successfully")

	var resp dto.ChapterResponse
	resp.FromModel(existingChapter)
	return &resp, nil
}

// DeleteChapter deletes a chapter by ID
func (s *chapterService) DeleteChapter(ctx context.Context, id uint) error {
	// Check if chapter exists
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("chapter_id", id).Info("Chapter not found for deletion")
			return fmt.Errorf("chapter with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("chapter_id", id).Error("Error finding chapter for deletion")
		return err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.WithError(err).WithField("chapter_id", id).Error("Failed to delete chapter")
		return err
	}

	s.logger.WithField("chapter_id", id).Info("Chapter deleted successfully")
	return nil
}

// GetChapterByID retrieves a chapter by ID
func (s *chapterService) GetChapterByID(ctx context.Context, id uint) (*dto.ChapterResponse, error) {
	chapter, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("chapter_id", id).Info("Chapter not found")
			return nil, fmt.Errorf("chapter with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("chapter_id", id).Error("Error finding chapter")
		return nil, err
	}

	var resp dto.ChapterResponse
	resp.FromModel(chapter)
	return &resp, nil
}

// GetChaptersByCourseID retrieves all chapters for a specific course
func (s *chapterService) GetChaptersByCourseID(ctx context.Context, courseID uint) ([]dto.ChapterResponse, error) {
	// Verify course exists
	_, err := s.courseRepo.FindByID(ctx, courseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("course_id", courseID).Info("Course not found for chapter retrieval")
			return nil, fmt.Errorf("course with ID %d not found", courseID)
		}
		s.logger.WithError(err).WithField("course_id", courseID).Error("Error finding course for chapter retrieval")
		return nil, err
	}

	chapters, err := s.repo.FindByCourseID(ctx, courseID)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", courseID).Error("Failed to retrieve chapters for course")
		return nil, err
	}

	responses := make([]dto.ChapterResponse, len(chapters))
	for i, chapter := range chapters {
		responses[i].FromModel(&chapter)
	}
	return responses, nil
}

// GetChapterWithLessons retrieves a chapter with its lessons
func (s *chapterService) GetChapterWithLessons(ctx context.Context, id uint) (*dto.ChapterDetailResponse, error) {
	chapter, err := s.repo.FindByIDWithLessons(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("chapter_id", id).Info("Chapter not found")
			return nil, fmt.Errorf("chapter with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("chapter_id", id).Error("Error finding chapter with lessons")
		return nil, err
	}

	var resp dto.ChapterDetailResponse
	resp.FromModel(chapter)
	return &resp, nil
}
