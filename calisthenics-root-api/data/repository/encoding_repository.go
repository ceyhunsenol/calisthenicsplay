package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IEncodingRepository interface {
	Save(encoding data.Encoding) (*data.Encoding, error)
	GetAll() ([]data.Encoding, error)
	GetByID(id string) (*data.Encoding, error)
	Update(encoding data.Encoding) (*data.Encoding, error)
	Delete(id string) error
}

type encodingRepository struct {
	DB *gorm.DB
}

func NewEncodingRepository(db *gorm.DB) IEncodingRepository {
	return &encodingRepository{DB: db}
}

func (r *encodingRepository) Save(encoding data.Encoding) (*data.Encoding, error) {
	result := r.DB.Create(&encoding)
	return &encoding, result.Error
}

func (r *encodingRepository) GetAll() ([]data.Encoding, error) {
	var encodings []data.Encoding
	result := r.DB.Find(&encodings)
	return encodings, result.Error
}

func (r *encodingRepository) GetByID(id string) (*data.Encoding, error) {
	var encoding data.Encoding
	result := r.DB.Where("id = ?", id).First(&encoding)
	return &encoding, result.Error
}

func (r *encodingRepository) Update(encoding data.Encoding) (*data.Encoding, error) {
	result := r.DB.Updates(&encoding)
	return &encoding, result.Error
}

func (r *encodingRepository) Delete(id string) error {
	return r.DB.Delete(&data.Encoding{}, "id = ?", id).Error
}
