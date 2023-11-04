package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IGenreContentRepository interface {
	Save(genreContent data.GenreContent) (*data.GenreContent, error)
	Delete(genreContent data.GenreContent) error
}

type genreContentRepository struct {
	DB *gorm.DB
}

func NewGenreContentRepository(db *gorm.DB) IGenreContentRepository {
	return &genreContentRepository{DB: db}
}

func (r *genreContentRepository) Save(genreContent data.GenreContent) (*data.GenreContent, error) {
	result := r.DB.Create(&genreContent)
	return &genreContent, result.Error
}

func (r *genreContentRepository) Delete(genreContent data.GenreContent) error {
	return r.DB.Delete(&data.GenreContent{}, "genre_id = ? AND content_id = ?", genreContent.GenreID, genreContent.ContentID).Error
}
