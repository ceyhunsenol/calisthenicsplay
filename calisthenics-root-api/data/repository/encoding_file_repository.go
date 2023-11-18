package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IEncodingFileRepository interface {
	Save(encodingFile data.EncodingFile) (*data.EncodingFile, error)
	GetAllEncodingFilesByEncodingID(encodingID string) ([]data.EncodingFile, error)
	GetByID(id string) (*data.EncodingFile, error)
	Update(encodingFile data.EncodingFile) (*data.EncodingFile, error)
	Delete(id string) error
	DeleteEncodingFilesByEncodingID(id string) error
}

type encodingFileRepository struct {
	DB *gorm.DB
}

func NewEncodingFileRepository(db *gorm.DB) IEncodingFileRepository {
	return &encodingFileRepository{DB: db}
}

func (r *encodingFileRepository) Save(encodingFile data.EncodingFile) (*data.EncodingFile, error) {
	result := r.DB.Create(&encodingFile)
	return &encodingFile, result.Error
}

func (r *encodingFileRepository) GetAllEncodingFilesByEncodingID(encodingID string) ([]data.EncodingFile, error) {
	var encodingFiles []data.EncodingFile
	result := r.DB.Where("encoding_id = ?", encodingID).Find(&encodingFiles)
	return encodingFiles, result.Error
}

func (r *encodingFileRepository) GetByID(id string) (*data.EncodingFile, error) {
	var encodingFile data.EncodingFile
	result := r.DB.Where("id = ?", id).First(&encodingFile)
	return &encodingFile, result.Error
}

func (r *encodingFileRepository) Update(encodingFile data.EncodingFile) (*data.EncodingFile, error) {
	result := r.DB.Updates(&encodingFile)
	return &encodingFile, result.Error
}

func (r *encodingFileRepository) Delete(id string) error {
	return r.DB.Delete(&data.EncodingFile{}, "id = ?", id).Error
}

func (r *encodingFileRepository) DeleteEncodingFilesByEncodingID(id string) error {
	return r.DB.Delete(&data.EncodingFile{}, "encoding_id = ?", id).Error
}
