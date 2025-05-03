package mapper

import (
	"LmsSystem/dto"
	"LmsSystem/models"
)

func ToLessonDTO(model models.Lesson) dto.LessonDTO {
	return dto.LessonDTO{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Content:     model.Content,
		Order:       model.Order,
		ChapterID:   model.ChapterID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func ToLessonModel(dto dto.LessonDTO) models.Lesson {
	return models.Lesson{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Content:     dto.Content,
		Order:       dto.Order,
		ChapterID:   dto.ChapterID,
	}
}
