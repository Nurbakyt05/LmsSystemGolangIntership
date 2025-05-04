package handler

import (
	"LmsSystem/dto"
	"LmsSystem/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// CourseHandler handles HTTP requests for courses
type CourseHandler struct {
	svc    service.CourseService
	logger *logrus.Logger
}

// NewCourseHandler creates a new CourseHandler
func NewCourseHandler(svc service.CourseService, logger *logrus.Logger) *CourseHandler {
	return &CourseHandler{svc: svc, logger: logger}
}

// GetAllCourses returns all courses
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.svc.GetAllCourses(c.Request.Context())
	if err != nil {
		h.logger.WithError(err).Error("Failed to retrieve courses")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

// CreateCourse creates a new course
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req dto.CourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid create course request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	resp, err := h.svc.CreateCourse(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create course")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetCourseByID returns a course by its ID
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("course_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}
	course, err := h.svc.GetCourseByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to retrieve course")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

// UpdateCourse updates an existing course
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("course_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}
	var req dto.CourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid update course request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	resp, err := h.svc.UpdateCourse(c.Request.Context(), uint(id), &req)
	if err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to update course")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found or update failed"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCourse deletes a course by its ID
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("course_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}
	if err := h.svc.DeleteCourse(c.Request.Context(), uint(id)); err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to delete course")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found or deletion failed"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetCourseWithChapters returns a course with its chapters
func (h *CourseHandler) GetCourseWithChapters(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("course_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}
	detail, err := h.svc.GetCourseWithChapters(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField(`course_id`, id).Error("Failed to retrieve course with chapters")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, detail)
}
