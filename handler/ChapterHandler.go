package handler

import (
	"LmsSystem/dto"
	"LmsSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ChapterHandler struct {
	service service.ChapterService
	logger  *logrus.Logger
}

func NewChapterHandler(s service.ChapterService, logger *logrus.Logger) *ChapterHandler {
	return &ChapterHandler{service: s, logger: logger}
}

func (h *ChapterHandler) CreateChapter(c *gin.Context) {
	var req dto.ChapterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	resp, err := h.service.CreateChapter(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create chapter")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *ChapterHandler) GetChapterByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("chapter_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid chapter ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID format"})
		return
	}
	resp, err := h.service.GetChapterByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).WithField("chapter_id", id).Error("Chapter not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *ChapterHandler) GetChaptersByCourseID(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("course_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid course ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID format"})
		return
	}
	list, err := h.service.GetChaptersByCourseID(c.Request.Context(), uint(courseID))
	if err != nil {
		h.logger.WithError(err).WithField("course_id", courseID).Error("Failed to retrieve chapters")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve chapters"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *ChapterHandler) GetChapterWithLessons(c *gin.Context) {
	rawID := c.Param("chapter_id")
	id64, err := strconv.ParseUint(rawID, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid chapter ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID format"})
		return
	}

	detail, err := h.service.GetChapterWithLessons(c.Request.Context(), uint(id64))
	if err != nil {
		h.logger.WithError(err).WithField("chapter_id", id64).Error("Failed to retrieve chapter with lessons")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

func (h *ChapterHandler) UpdateChapter(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("chapter_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid chapter ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID format"})
		return
	}
	var req dto.ChapterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	updated, err := h.service.UpdateChapter(c.Request.Context(), uint(id), &req)
	if err != nil {
		h.logger.WithError(err).WithField("chapter_id", id).Error("Failed to update chapter")
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found or update failed"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ChapterHandler) DeleteChapter(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("chapter_id"), 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid chapter ID format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID format"})
		return
	}
	if err := h.service.DeleteChapter(c.Request.Context(), uint(id)); err != nil {
		h.logger.WithError(err).WithField("chapter_id", id).Error("Failed to delete chapter")
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found or deletion failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
