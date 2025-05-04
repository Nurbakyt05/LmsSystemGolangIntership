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

// CourseService defines the interface for course business logic
type CourseService interface {
	CreateCourse(ctx context.Context, courseReq *dto.CourseRequest) (*dto.CourseResponse, error)
	UpdateCourse(ctx context.Context, id uint, courseReq *dto.CourseRequest) (*dto.CourseResponse, error)
	DeleteCourse(ctx context.Context, id uint) error
	GetCourseByID(ctx context.Context, id uint) (*dto.CourseResponse, error)
	GetAllCourses(ctx context.Context) ([]dto.CourseResponse, error)
	GetCourseWithChapters(ctx context.Context, id uint) (*dto.CourseDetailResponse, error)
}

// courseService implements CourseService interface
type courseService struct {
	repo   repository.CourseRepository
	logger *logrus.Logger
}

// NewCourseService creates a new instance of CourseService
func NewCourseService(repo repository.CourseRepository, logger *logrus.Logger) CourseService {
	return &courseService{
		repo:   repo,
		logger: logger,
	}
}

// CreateCourse creates a new course
func (s *courseService) CreateCourse(ctx context.Context, courseReq *dto.CourseRequest) (*dto.CourseResponse, error) {
	course := courseReq.ToModel()
	err := s.repo.Create(ctx, course)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create course")
		return nil, err
	}

	s.logger.WithFields(logrus.Fields{
		"course_id":   course.ID,
		"course_name": course.Name,
	}).Info("Course created successfully")

	var resp dto.CourseResponse
	resp.FromModel(course)
	return &resp, nil
}

// UpdateCourse updates an existing course
func (s *courseService) UpdateCourse(ctx context.Context, id uint, courseReq *dto.CourseRequest) (*dto.CourseResponse, error) {
	existingCourse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("course_id", id).Info("Course not found for update")
			return nil, fmt.Errorf("course with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("course_id", id).Error("Error finding course for update")
		return nil, err
	}

	existingCourse.Name = courseReq.Name
	existingCourse.Description = courseReq.Description

	err = s.repo.Update(ctx, existingCourse)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", id).Error("Failed to update course")
		return nil, err
	}

	s.logger.WithField("course_id", id).Info("Course updated successfully")

	var resp dto.CourseResponse
	resp.FromModel(existingCourse)
	return &resp, nil
}

// DeleteCourse deletes a course by ID
func (s *courseService) DeleteCourse(ctx context.Context, id uint) error {
	// Check if course exists
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("course_id", id).Info("Course not found for deletion")
			return fmt.Errorf("course with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("course_id", id).Error("Error finding course for deletion")
		return err
	}

	err = s.repo.Delete(ctx, id)
	if err != nil {
		s.logger.WithError(err).WithField("course_id", id).Error("Failed to delete course")
		return err
	}

	s.logger.WithField("course_id", id).Info("Course deleted successfully")
	return nil
}

// GetCourseByID retrieves a course by ID
func (s *courseService) GetCourseByID(ctx context.Context, id uint) (*dto.CourseResponse, error) {
	course, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("course_id", id).Info("Course not found")
			return nil, fmt.Errorf("course with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("course_id", id).Error("Error finding course")
		return nil, err
	}

	var resp dto.CourseResponse
	resp.FromModel(course)
	return &resp, nil
}

// GetAllCourses retrieves all courses
func (s *courseService) GetAllCourses(ctx context.Context) ([]dto.CourseResponse, error) {
	courses, err := s.repo.FindAll(ctx)
	if err != nil {
		s.logger.WithError(err).Error("Failed to retrieve all courses")
		return nil, err
	}

	responses := make([]dto.CourseResponse, len(courses))
	for i, course := range courses {
		responses[i].FromModel(&course)
	}
	return responses, nil
}

// GetCourseWithChapters retrieves a course with its chapters
func (s *courseService) GetCourseWithChapters(ctx context.Context, id uint) (*dto.CourseDetailResponse, error) {
	course, err := s.repo.FindByIDWithChapters(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.WithField("course_id", id).Info("Course not found")
			return nil, fmt.Errorf("course with ID %d not found", id)
		}
		s.logger.WithError(err).WithField("course_id", id).Error("Error finding course with chapters")
		return nil, err
	}

	var resp dto.CourseDetailResponse
	resp.FromModel(course)
	return &resp, nil
}
