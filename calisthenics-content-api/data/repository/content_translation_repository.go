package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type IContentTranslationRepository interface {
	GetAll() ([]data.ContentTranslation, error)
	GetByID(id string) (*data.ContentTranslation, error)
	GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error)
	GetAllByCode(code string) ([]data.ContentTranslation, error)
}

type contentTranslationRepository struct {
	DB *gorm.DB
}

func NewContentTranslationRepository(db *gorm.DB) IContentTranslationRepository {
	return &contentTranslationRepository{DB: db}
}

func (r *contentTranslationRepository) GetAll() ([]data.ContentTranslation, error) {
	var translations []data.ContentTranslation
	result := r.DB.Find(&translations)
	return translations, result.Error
}

func (r *contentTranslationRepository) GetByID(id string) (*data.ContentTranslation, error) {
	var translation data.ContentTranslation
	result := r.DB.Where("id = ?", id).First(&translation)
	return &translation, result.Error
}

func (r *contentTranslationRepository) GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error) {
	var translation data.ContentTranslation
	result := r.DB.Where("code = ? and lang_code = ?", code, langCode).First(&translation)
	return &translation, result.Error
}

func (r *contentTranslationRepository) GetAllByCode(code string) ([]data.ContentTranslation, error) {
	var translations []data.ContentTranslation
	result := r.DB.Where("code = ?", code).Find(&translations)
	return translations, result.Error
}
