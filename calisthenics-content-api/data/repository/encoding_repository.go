package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type IEncodingRepository interface {
	GetAll() ([]data.Encoding, error)
	GetByID(id string) (*data.Encoding, error)
}

type encodingRepository struct {
	DB *gorm.DB
}

func NewEncodingRepository(db *gorm.DB) IEncodingRepository {
	return &encodingRepository{DB: db}
}

func (r *encodingRepository) GetAll() ([]data.Encoding, error) {
	var encodings []data.Encoding
	result := r.DB.Preload("EncodingFiles").Find(&encodings)
	return encodings, result.Error
}

func (r *encodingRepository) GetByID(id string) (*data.Encoding, error) {
	var encoding data.Encoding
	result := r.DB.Where("id = ?", id).First(&encoding)
	return &encoding, result.Error
}
