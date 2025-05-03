package routes

import (
	"LmsSystem/handler"
	"LmsSystem/repository"
	"LmsSystem/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")

	// ===== Course =====
	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	api.GET("/get-courses", courseHandler.GetAll)
	api.GET("/get-course/:id", courseHandler.GetByID)
	api.POST("/create-course", courseHandler.Create)
	api.PUT("/update-course/:id", courseHandler.Update)
	api.DELETE("/delete-course/:id", courseHandler.Delete)

	// ===== Chapter =====
	chapterRepo := repository.NewChapterRepository(db)
	chapterService := service.NewChapterService(chapterRepo)
	chapterHandler := handler.NewChapterHandler(chapterService)

	api.GET("/get-chapters", chapterHandler.GetAll)
	api.GET("/get-chapter/:id", chapterHandler.GetByID)
	api.POST("/create-chapter", chapterHandler.Create)
	api.PUT("/update-chapter/:id", chapterHandler.Update)
	api.DELETE("/delete-chapter/:id", chapterHandler.Delete)

	// ===== Lesson =====
	lessonRepo := repository.NewLessonRepository(db)
	lessonService := service.NewLessonService(lessonRepo)
	lessonHandler := handler.NewLessonHandler(lessonService)

	api.GET("/get-lessons", lessonHandler.GetAll)
	api.GET("/get-lesson/:id", lessonHandler.GetByID)
	api.POST("/create-lesson", lessonHandler.Create)
	api.PUT("/update-lesson/:id", lessonHandler.Update)
	api.DELETE("/delete-lesson/:id", lessonHandler.Delete)
}
