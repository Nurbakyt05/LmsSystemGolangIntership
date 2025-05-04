package router

import (
	"LmsSystem/handler"
	"LmsSystem/repository"
	"LmsSystem/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, r *gin.Engine) {
	// Репозитории
	courseRepo := repository.NewCourseRepository(db)
	chapterRepo := repository.NewChapterRepository(db)
	lessonRepo := repository.NewLessonRepository(db)

	// Сервисы
	courseService := service.NewCourseService(courseRepo)
	chapterService := service.NewChapterService(chapterRepo)
	lessonService := service.NewLessonService(lessonRepo)

	// Хендлеры
	courseHandler := handler.NewCourseHandler(courseService)
	chapterHandler := handler.NewChapterHandler(chapterService)
	lessonHandler := handler.NewLessonHandler(lessonService)

	api := r.Group("/api")

	// ======= Курсы =======
	courses := api.Group("/courses")
	{
		courses.GET("/", courseHandler.GetAll)
		courses.GET("/full", courseHandler.GetAllFullCourses) // все курсы + их главы и уроки
		courses.POST("/", courseHandler.Create)

		courseItem := courses.Group("/:id")
		{
			courseItem.GET("", courseHandler.GetByID)
			courseItem.GET("/full", courseHandler.GetFullCourse) // один курс + его главы и уроки
			courseItem.PUT("", courseHandler.Update)
			courseItem.DELETE("", courseHandler.Delete)
			courseItem.GET("/chapters", chapterHandler.GetChaptersByCourse)
		}
	}

	// ======= Главы =======
	chapters := api.Group("/chapters")
	{
		chapters.GET("/full", chapterHandler.GetAllFullChapters) // все главы + уроки
		chapters.GET("/", chapterHandler.GetAllChapters)
		chapters.POST("/", chapterHandler.CreateChapter)

		chapterItem := chapters.Group("/:id")
		{
			chapterItem.GET("/full", chapterHandler.GetFullChapterByID) // одна глава + уроки
			chapterItem.GET("", chapterHandler.GetChapter)
			chapterItem.PUT("", chapterHandler.UpdateChapter)
			chapterItem.DELETE("", chapterHandler.DeleteChapter)
			chapterItem.GET("/lessons", lessonHandler.GetLessonsByChapter)
		}
	}

	// ======= Уроки =======
	lessons := api.Group("/lessons")
	{
		lessons.GET("/full", lessonHandler.GetAllFullLessons) // все уроки + их глава
		lessons.GET("/", lessonHandler.GetAllLessons)
		lessons.POST("/", lessonHandler.CreateLesson)

		lessonItem := lessons.Group("/:id")
		{
			lessonItem.GET("/full", lessonHandler.GetFullLessonByID) // один урок + его глава
			lessonItem.GET("", lessonHandler.GetLesson)
			lessonItem.PUT("", lessonHandler.UpdateLesson)
			lessonItem.DELETE("", lessonHandler.DeleteLesson)
		}
	}
}
