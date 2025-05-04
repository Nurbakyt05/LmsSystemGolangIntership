package handler

import (
	"LmsSystem/models"
	"LmsSystem/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChapterHandler struct {
	service service.ChapterService
}

func NewChapterHandler(service service.ChapterService) *ChapterHandler {
	return &ChapterHandler{service: service}
}

func (h *ChapterHandler) GetAllChapters(c *gin.Context) {
	chapters, err := h.service.GetAllChapters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chapters)
}

func (h *ChapterHandler) GetChapter(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	chapter, err := h.service.GetChapterByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}
	c.JSON(http.StatusOK, chapter)
}

func (h *ChapterHandler) GetChaptersByCourse(c *gin.Context) {
	courseID, err := strconv.Atoi(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	chapters, err := h.service.GetChaptersByCourse(uint(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chapters)
}

func (h *ChapterHandler) CreateChapter(c *gin.Context) {
	var chapter models.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateChapter(&chapter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, chapter)
}

func (h *ChapterHandler) UpdateChapter(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	var chapter models.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	chapter.ID = uint(id)

	if err := h.service.UpdateChapter(&chapter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, chapter)
}

func (h *ChapterHandler) DeleteChapter(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	if err := h.service.DeleteChapter(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Chapter deleted successfully"})
}
