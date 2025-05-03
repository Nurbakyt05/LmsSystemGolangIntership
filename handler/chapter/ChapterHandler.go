package chapter

import (
	dto "LmsSystem/dto/chapter"
	"LmsSystem/service/chapter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ChapterHandler struct {
	service chapter.ChapterService
}

func NewChapterHandler(service chapter.ChapterService) *ChapterHandler {
	return &ChapterHandler{service: service}
}

func (h *ChapterHandler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/chapters")
	{
		group.GET("", h.GetAll)
		group.GET("/:id", h.GetByID)
		group.POST("", h.Create)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
	}
}

func (h *ChapterHandler) GetAll(c *gin.Context) {
	chapters, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch chapters"})
		return
	}
	c.JSON(http.StatusOK, chapters)
}

func (h *ChapterHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	chapterDTO, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Chapter not found"})
		return
	}
	c.JSON(http.StatusOK, chapterDTO)
}

func (h *ChapterHandler) Create(c *gin.Context) {
	var input dto.ChapterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chapter"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Chapter created"})
}

func (h *ChapterHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	var input dto.ChapterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	input.ID = uint(id)

	if err := h.service.Update(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update chapter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Chapter updated"})
}

func (h *ChapterHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete chapter"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Chapter deleted"})
}
