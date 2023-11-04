package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IMediaRepository interface {
	Save(media data.Media) (*data.Media, error)
	GetAll() ([]data.Media, error)
	GetByID(id string) (*data.Media, error)
	Update(media data.Media) (*data.Media, error)
	Delete(id string) error
}

type mediaRepository struct {
	DB *gorm.DB
}

func NewMediaRepository(db *gorm.DB) IMediaRepository {
	return &mediaRepository{DB: db}
}

func (r *mediaRepository) Save(media data.Media) (*data.Media, error) {
	result := r.DB.Create(&media)
	return &media, result.Error
}

func (r *mediaRepository) GetAll() ([]data.Media, error) {
	var medias []data.Media
	result := r.DB.Find(&medias)
	return medias, result.Error
}

func (r *mediaRepository) GetByID(id string) (*data.Media, error) {
	var media data.Media
	result := r.DB.Where("id = ?", id).First(&media)
	return &media, result.Error
}

func (r *mediaRepository) Update(media data.Media) (*data.Media, error) {
	result := r.DB.Updates(&media)
	if result.Error != nil {
		return nil, result.Error
	}
	return &media, nil
}

func (r *mediaRepository) Delete(id string) error {
	return r.DB.Delete(&data.Media{}, "id = ?", id).Error
}
