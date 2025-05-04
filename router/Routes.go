package router

import (
	"net/http"

	"LmsSystem/handler"
	"LmsSystem/repository"
	"LmsSystem/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, r *gin.Engine, logger *logrus.Logger) {
	// Repositories
	courseRepo := repository.NewCourseRepository(db)
	chapterRepo := repository.NewChapterRepository(db)
	lessonRepo := repository.NewLessonRepository(db)

	// Services
	courseSvc := service.NewCourseService(courseRepo, logger)
	chapterSvc := service.NewChapterService(chapterRepo, courseRepo, logger)
	lessonSvc := service.NewLessonService(lessonRepo, chapterRepo, logger)

	// Handlers
	courseH := handler.NewCourseHandler(courseSvc, logger)
	chapterH := handler.NewChapterHandler(chapterSvc, logger)
	lessonH := handler.NewLessonHandler(lessonSvc, logger)

	api := r.Group("/api")
	{
		// Health-check
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// Courses
		courses := api.Group("/courses")
		{
			courses.GET("", courseH.GetAllCourses)
			courses.POST("", courseH.CreateCourse)

			byCourse := courses.Group("/:course_id")
			{
				byCourse.GET("", courseH.GetCourseByID)
				byCourse.PUT("", courseH.UpdateCourse)
				byCourse.DELETE("", courseH.DeleteCourse)
				byCourse.GET("/chapters", chapterH.GetChaptersByCourseID)
			}
		}

		// Chapters
		chapters := api.Group("/chapters")
		{
			chapters.POST("", chapterH.CreateChapter)

			byChapter := chapters.Group("/:chapter_id")
			{
				// Возвращает главу вместе со всеми уроками
				byChapter.GET("", chapterH.GetChapterWithLessons)
				byChapter.PUT("", chapterH.UpdateChapter)
				byChapter.DELETE("", chapterH.DeleteChapter)
				// убрали: byChapter.GET("/lessons", lessonH.GetLessonsByChapterID)
			}
		}

		// Lessons
		lessons := api.Group("/lessons")
		{
			lessons.POST("", lessonH.CreateLesson)

			byLesson := lessons.Group("/:lesson_id")
			{
				byLesson.GET("", lessonH.GetLessonByID)
				byLesson.PUT("", lessonH.UpdateLesson)
				byLesson.DELETE("", lessonH.DeleteLesson)
			}
		}
	}
}
