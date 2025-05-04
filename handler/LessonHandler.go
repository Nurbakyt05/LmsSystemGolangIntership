package handler

import (
	"LmsSystem/dto"
	"LmsSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LessonHandler struct {
	service service.LessonService
	logger  *logrus.Logger
}

func NewLessonHandler(s service.LessonService, logger *logrus.Logger) *LessonHandler {
	return &LessonHandler{service: s, logger: logger}
}

// RegisterRoutes registers lesson endpoints
func (h *LessonHandler) RegisterRoutes(r *gin.RouterGroup) {
	ls := r.Group("/lessons")
	{
		ls.POST("", h.CreateLesson)
		ls.GET("/:id", h.GetLessonByID)
		ls.PUT("/:id", h.UpdateLesson)
		ls.DELETE("/:id", h.DeleteLesson)
	}
	r.GET("/chapters/:chapterId/lessons", h.GetLessonsByChapterID)
}

func (h *LessonHandler) CreateLesson(c *gin.Context) {
	var req dto.LessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid lesson creation request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	resp, err := h.service.CreateLesson(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create lesson")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lesson"})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *LessonHandler) GetLessonByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid lesson ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID format"})
		return
	}
	resp, err := h.service.GetLessonByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField("lesson_id", id).Error("Lesson not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LessonHandler) GetLessonsByChapterID(c *gin.Context) {
	chapterID, err := strconv.ParseUint(c.Param("chapterId"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid chapter ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID format"})
		return
	}
	list, err := h.service.GetLessonsByChapterID(c.Request.Context(), uint(chapterID))
	if err != nil {
		h.logger.WithError(err).WithField("chapter_id", chapterID).Error("Failed to retrieve lessons")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lessons"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *LessonHandler) UpdateLesson(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid lesson ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID format"})
		return
	}
	var req dto.LessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	resp, err := h.service.UpdateLesson(c.Request.Context(), uint(id), &req)
	if err != nil {
		h.logger.WithError(err).WithField("lesson_id", id).Error("Failed to update lesson")
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found or update failed"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LessonHandler) DeleteLesson(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid lesson ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID format"})
		return
	}
	if err := h.service.DeleteLesson(c.Request.Context(), uint(id)); err != nil {
		h.logger.WithError(err).WithField("lesson_id", id).Error("Failed to delete lesson")
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found or deletion failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
