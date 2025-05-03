package chapter

import (
	dto "LmsSystem/dto/chapter"
	"LmsSystem/models"
)

func ToChapterDTO(chapter models.Chapter) dto.ChapterDTO {
	return dto.ChapterDTO{
		ID:          chapter.ID,
		Name:        chapter.Name,
		Description: chapter.Description,
		Order:       chapter.Order,
		CourseID:    chapter.CourseID,
		CreatedAt:   chapter.CreatedAt,
		UpdatedAt:   chapter.UpdatedAt,
	}
}

func ToChapterModel(dto dto.ChapterDTO) models.Chapter {
	return models.Chapter{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Order:       dto.Order,
		CourseID:    dto.CourseID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}
