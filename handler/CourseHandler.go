package handler

import (
	"LmsSystem/dto"
	"LmsSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CourseHandler handles HTTP requests related to courses
type CourseHandler struct {
	courseService service.CourseService
	logger        *logrus.Logger
}

// NewCourseHandler creates a new instance of CourseHandler
func NewCourseHandler(courseService service.CourseService, logger *logrus.Logger) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
		logger:        logger,
	}
}

// RegisterRoutes registers course routes to the given router group
func (h *CourseHandler) RegisterRoutes(router *gin.RouterGroup) {
	courses := router.Group("/courses")
	{
		courses.POST("", h.CreateCourse)
		courses.GET("", h.GetAllCourses)
		courses.GET("/:id", h.GetCourseByID)
		courses.GET("/:id/chapters", h.GetCourseWithChapters)
		courses.PUT("/:id", h.UpdateCourse)
		courses.DELETE("/:id", h.DeleteCourse)
	}
}

// CreateCourse handles the creation of a new course
// @Summary Create a new course
// @Description Create a new course with the provided information
// @Tags courses
// @Accept json
// @Produce json
// @Param course body dto.CourseRequest true "Course information"
// @Success 201 {object} dto.CourseResponse "Created course"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/courses [post]
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var courseReq dto.CourseRequest
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		h.logger.WithError(err).Error("Invalid course creation request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.courseService.CreateCourse(c.Request.Context(), &courseReq)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create course")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetAllCourses handles the retrieval of all courses
// @Summary Get all courses
// @Description Get a list of all courses
// @Tags courses
// @Produce json
// @Success 200 {array} dto.CourseResponse "List of courses"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/courses [get]
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.courseService.GetAllCourses(c.Request.Context())
	if err != nil {
		h.logger.WithError(err).Error("Failed to retrieve courses")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// GetCourseByID handles the retrieval of a course by ID
// @Summary Get a course by ID
// @Description Get a course by its ID
// @Tags courses
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} dto.CourseResponse "Course details"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Course not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/courses/{id} [get]
func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}

	course, err := h.courseService.GetCourseByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to retrieve course")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// GetCourseWithChapters handles the retrieval of a course with its chapters
// @Summary Get a course with chapters
// @Description Get a course with its chapters by course ID
// @Tags courses
// @Produce json
// @Param id path int true "Course ID"
// @Success 200 {object} dto.CourseDetailResponse "Course with chapters"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Course not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/courses/{id}/chapters [get]
func (h *CourseHandler) GetCourseWithChapters(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}

	course, err := h.courseService.GetCourseWithChapters(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to retrieve course with chapters")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// UpdateCourse handles the update of an existing course
// @Summary Update a course
// @Description Update a course with the provided information
// @Tags courses
// @Accept json
// @Produce json
// @Param id path int true "Course ID"
// @Param course body dto.CourseRequest true "Updated course information"
// @Success 200 {object} dto.CourseResponse "Updated course"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Course not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/courses/{id} [put]
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}

	var courseReq dto.CourseRequest
	if err := c.ShouldBindJSON(&courseReq); err != nil {
		h.logger.WithError(err).Error("Invalid course update request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.courseService.UpdateCourse(c.Request.Context(), uint(id), &courseReq)
	if err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to update course")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteCourse handles the deletion of a course
// @Summary Delete a course
// @Description Delete a course by its ID
// @Tags courses
// @Param id path int true "Course ID"
// @Success 204 "No content"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Course not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/courses/{id} [delete]
func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}

	err = h.courseService.DeleteCourse(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField("course_id", id).Error("Failed to delete course")
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found or deletion failed"})
		return
	}

	c.Status(http.StatusNoContent)
}
