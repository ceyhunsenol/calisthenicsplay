package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IGenreTypeRepository interface {
	Save(genreType data.GenreType) (*data.GenreType, error)
	GetAll() ([]data.GenreType, error)
	GetByID(id string) (*data.GenreType, error)
	GetByCode(code string) (*data.GenreType, error)
	ExistsByCode(code string) (bool, error)
	Update(genreType data.GenreType) (*data.GenreType, error)
	Delete(id string) error
}

type genreTypeRepository struct {
	DB *gorm.DB
}

func NewGenreTypeRepository(db *gorm.DB) IGenreTypeRepository {
	return &genreTypeRepository{DB: db}
}

func (r *genreTypeRepository) Save(genreType data.GenreType) (*data.GenreType, error) {
	result := r.DB.Create(&genreType)
	return &genreType, result.Error
}

func (r *genreTypeRepository) GetAll() ([]data.GenreType, error) {
	var genreTypes []data.GenreType
	result := r.DB.Find(&genreTypes)
	return genreTypes, result.Error
}

func (r *genreTypeRepository) GetByID(id string) (*data.GenreType, error) {
	var genreType data.GenreType
	result := r.DB.Where("id = ?", id).First(&genreType)
	return &genreType, result.Error
}

func (r *genreTypeRepository) GetByCode(code string) (*data.GenreType, error) {
	var genreType data.GenreType
	result := r.DB.Where("code = ?", code).First(&genreType)
	return &genreType, result.Error
}

func (r *genreTypeRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.GenreType{}).Where("code = ?", code).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *genreTypeRepository) Update(genreType data.GenreType) (*data.GenreType, error) {
	result := r.DB.Updates(&genreType)
	if result.Error != nil {
		return nil, result.Error
	}
	return &genreType, nil
}

func (r *genreTypeRepository) Delete(id string) error {
	return r.DB.Delete(&data.GenreType{}, "id = ?", id).Error
}
