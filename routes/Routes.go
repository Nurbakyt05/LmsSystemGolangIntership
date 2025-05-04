package router

import (
	"LmsSystem/handler"
	"LmsSystem/repository"
	"LmsSystem/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Инициализация репозиториев
	courseRepo := repository.NewCourseRepository(db)
	chapterRepo := repository.NewChapterRepository(db)
	lessonRepo := repository.NewLessonRepository(db)

	// Инициализация сервисов
	courseService := service.NewCourseService(courseRepo)
	chapterService := service.NewChapterService(chapterRepo)
	lessonService := service.NewLessonService(lessonRepo)

	// Инициализация хендлеров
	courseHandler := handler.NewCourseHandler(courseService)
	chapterHandler := handler.NewChapterHandler(chapterService)
	lessonHandler := handler.NewLessonHandler(lessonService)

	// Маршруты API
	api := r.Group("/api")
	{
		// Курсы
		courses := api.Group("/courses")
		{
			courses.GET("/", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)
			courses.GET("/", courseHandler.GetFullCourse)
			courses.POST("/", courseHandler.Create)
			courses.PUT("/:id", courseHandler.Update)
			courses.DELETE("/:id", courseHandler.Delete)

			// Главы курса
			courses.GET("/:courseId/chapters", chapterHandler.GetChaptersByCourse)
		}

		// Главы
		chapters := api.Group("/chapters")
		{
			chapters.GET("/", chapterHandler.GetAllChapters)
			chapters.GET("/:id", chapterHandler.GetChapter)
			chapters.POST("/", chapterHandler.CreateChapter)
			chapters.PUT("/:id", chapterHandler.UpdateChapter)
			chapters.DELETE("/:id", chapterHandler.DeleteChapter)

			// Уроки главы
			chapters.GET("/:chapterId/lessons", lessonHandler.GetLessonsByChapter)
		}

		// Уроки
		lessons := api.Group("/lessons")
		{
			lessons.GET("/", lessonHandler.GetAllLessons)
			lessons.GET("/:id", lessonHandler.GetLesson)
			lessons.POST("/", lessonHandler.CreateLesson)
			lessons.PUT("/:id", lessonHandler.UpdateLesson)
			lessons.DELETE("/:id", lessonHandler.DeleteLesson)
		}
	}

	return r
}
