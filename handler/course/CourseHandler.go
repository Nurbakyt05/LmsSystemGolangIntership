package coursepackage

import (
	"LmsSystem/dto"
	service "LmsSystem/service/course"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(service service.CourseService) *CourseHandler {
	return &CourseHandler{service}
}

func (h *CourseHandler) RegisterRoutes(router *gin.RouterGroup) {
	courses := router.Group("/courses")
	{
		courses.GET("", h.GetAll)
		courses.GET("/:id", h.GetByID)
		courses.POST("", h.Create)
		courses.PUT("/:id", h.Update)
		courses.DELETE("/:id", h.Delete)
	}
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	course, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) Create(c *gin.Context) {
	var course dto.CourseDTO
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := h.service.Create(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create course"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Course created"})
}

func (h *CourseHandler) Update(c *gin.Context) {
	var course dto.CourseDTO
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	course.ID = uint(id)
	err := h.service.Update(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update course"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course updated"})
}

func (h *CourseHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete course"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted"})
}
