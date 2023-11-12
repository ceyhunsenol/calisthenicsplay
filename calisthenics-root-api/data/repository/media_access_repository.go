package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IMediaAccessRepository interface {
	Save(mediaAccess data.MediaAccess) (*data.MediaAccess, error)
	GetAll() ([]data.MediaAccess, error)
	GetByID(id string) (*data.MediaAccess, error)
	ExistsByMediaID(mediaID string) (bool, error)
	GetByMediaID(mediaID string) (*data.MediaAccess, error)
	Update(mediaAccess data.MediaAccess) (*data.MediaAccess, error)
	Delete(id string) error
	DeleteAllByMediaID(tx *gorm.DB, mediaID string) error
}

type mediaAccessRepository struct {
	DB *gorm.DB
}

func NewMediaAccessRepository(db *gorm.DB) IMediaAccessRepository {
	return &mediaAccessRepository{DB: db}
}

func (r *mediaAccessRepository) Save(mediaAccess data.MediaAccess) (*data.MediaAccess, error) {
	result := r.DB.Create(&mediaAccess)
	return &mediaAccess, result.Error
}

func (r *mediaAccessRepository) GetAll() ([]data.MediaAccess, error) {
	var mediaAccessList []data.MediaAccess
	result := r.DB.Find(&mediaAccessList)
	return mediaAccessList, result.Error
}

func (r *mediaAccessRepository) GetByID(id string) (*data.MediaAccess, error) {
	var mediaAccess data.MediaAccess
	result := r.DB.Where("id = ?", id).First(&mediaAccess)
	return &mediaAccess, result.Error
}

func (r *mediaAccessRepository) ExistsByMediaID(mediaID string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.MediaAccess{}).Where("media_id = ?", mediaID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *mediaAccessRepository) GetByMediaID(mediaID string) (*data.MediaAccess, error) {
	var mediaAccess data.MediaAccess
	result := r.DB.Where("media_id = ?", mediaID).First(&mediaAccess)
	return &mediaAccess, result.Error
}

func (r *mediaAccessRepository) Update(mediaAccess data.MediaAccess) (*data.MediaAccess, error) {
	result := r.DB.Updates(&mediaAccess)
	if result.Error != nil {
		return nil, result.Error
	}
	return &mediaAccess, nil
}

func (r *mediaAccessRepository) Delete(id string) error {
	return r.DB.Delete(&data.MediaAccess{}, "id = ?", id).Error
}

func (r *mediaAccessRepository) DeleteAllByMediaID(tx *gorm.DB, mediaID string) error {
	return tx.Delete(&data.MediaAccess{}, "media_id = ?", mediaID).Error
}
