package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type IContentRepository interface {
	GetAll() ([]data.Content, error)
	GetByID(id string) (*data.Content, error)
	GetByCode(code string) (*data.Content, error)
}

type contentRepository struct {
	DB *gorm.DB
}

func NewContentRepository(db *gorm.DB) IContentRepository {
	return &contentRepository{DB: db}
}

func (r *contentRepository) GetAll() ([]data.Content, error) {
	var contents []data.Content
	result := r.DB.Preload("Medias").Find(&contents)
	return contents, result.Error
}

func (r *contentRepository) GetByID(id string) (*data.Content, error) {
	var content data.Content
	result := r.DB.Where("id = ?", id).First(&content)
	return &content, result.Error
}

func (r *contentRepository) GetByCode(code string) (*data.Content, error) {
	var content data.Content
	result := r.DB.Where("code = ?", code).Preload("Translations").Find(&content)
	return &content, result.Error
}
