package mapper

import (
	"LmsSystem/dto"
	"LmsSystem/models"
)

func ToCourseDTO(course models.Course) dto.CourseDTO {
	return course.CourseDTO{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}
}

func ToCourseModel(courseDTO dto.CourseDTO) models.Course {
	return models.Course{
		ID:          courseDTO.ID,
		Name:        courseDTO.Name,
		Description: courseDTO.Description,
		CreatedAt:   courseDTO.CreatedAt,
		UpdatedAt:   courseDTO.UpdatedAt,
	}
}
