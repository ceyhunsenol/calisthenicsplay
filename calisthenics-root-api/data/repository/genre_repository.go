package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IGenreRepository interface {
	Save(tx *gorm.DB, genre data.Genre) (*data.Genre, error)
	GetAll() ([]data.Genre, error)
	GetByID(id string) (*data.Genre, error)
	GetByCode(code string) (*data.Genre, error)
	ExistsByCode(code string) (bool, error)
	Update(tx *gorm.DB, genre data.Genre) (*data.Genre, error)
	Delete(tx *gorm.DB, id string) error
}

type genreRepository struct {
	DB *gorm.DB
}

func NewGenreRepository(db *gorm.DB) IGenreRepository {
	return &genreRepository{DB: db}
}

func (r *genreRepository) Save(tx *gorm.DB, genre data.Genre) (*data.Genre, error) {
	result := tx.Create(&genre)
	return &genre, result.Error
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

func (r *genreRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.Genre{}).Where("code = ?", code).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *genreRepository) Update(tx *gorm.DB, genre data.Genre) (*data.Genre, error) {
	result := tx.Updates(&genre)
	if result.Error != nil {
		return nil, result.Error
	}
	return &genre, nil
}

func (r *genreRepository) Delete(tx *gorm.DB, id string) error {
	return tx.Delete(&data.Genre{}, "id = ?", id).Error
}
