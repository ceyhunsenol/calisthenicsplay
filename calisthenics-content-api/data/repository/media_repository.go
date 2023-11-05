package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type IMediaRepository interface {
	GetAll() ([]*data.Media, error)
	GetByID(id string) (*data.Media, error)
}

type mediaRepository struct {
	DB *gorm.DB
}

func NewMediaRepository(db *gorm.DB) IMediaRepository {
	return &mediaRepository{DB: db}
}

func (r *mediaRepository) GetAll() ([]*data.Media, error) {
	var medias []*data.Media
	result := r.DB.Find(&medias)
	return medias, result.Error
}

func (r *mediaRepository) GetByID(id string) (*data.Media, error) {
	var media data.Media
	result := r.DB.Where("id = ?", id).First(&media)
	return &media, result.Error
}
