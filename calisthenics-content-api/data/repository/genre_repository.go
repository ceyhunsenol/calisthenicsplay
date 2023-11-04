package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type IGenreRepository interface {
	GetAll() ([]data.Genre, error)
	GetByID(id string) (*data.Genre, error)
	GetByCode(code string) (*data.Genre, error)
}

type genreRepository struct {
	DB *gorm.DB
}

func NewGenreRepository(db *gorm.DB) IGenreRepository {
	return &genreRepository{DB: db}
}

func (r *genreRepository) GetAll() ([]data.Genre, error) {
	var genres []data.Genre
	result := r.DB.Preload("Contents").Find(&genres)
	return genres, result.Error
}

func (r *genreRepository) GetByID(id string) (*data.Genre, error) {
	var genre data.Genre
	result := r.DB.Where("id = ?", id).Preload("Contents").First(&genre)
	return &genre, result.Error
}

func (r *genreRepository) GetByCode(code string) (*data.Genre, error) {
	var genre data.Genre
	result := r.DB.Where("code = ?", code).First(&genre)
	return &genre, result.Error
}
