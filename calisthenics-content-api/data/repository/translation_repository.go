package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type ITranslationRepository interface {
	GetAll() ([]data.Translation, error)
	GetByID(id string) (*data.Translation, error)
	GetByCodeAndLangCode(id, langCode string) (*data.Translation, error)
	GetAllByCode(code string) ([]data.Translation, error)
	GetAllDistinctCodesByDomain(domain string) ([]string, error)
}

type translationRepository struct {
	DB *gorm.DB
}

func NewTranslationRepository(db *gorm.DB) ITranslationRepository {
	return &translationRepository{DB: db}
}

func (r *translationRepository) GetAll() ([]data.Translation, error) {
	var translations []data.Translation
	result := r.DB.Find(&translations)
	return translations, result.Error
}

func (r *translationRepository) GetByID(id string) (*data.Translation, error) {
	var translation data.Translation
	result := r.DB.Where("id = ?", id).First(&translation)
	return &translation, result.Error
}

func (r *translationRepository) GetByCodeAndLangCode(code, langCode string) (*data.Translation, error) {
	var translation data.Translation
	result := r.DB.Where("code = ? and lang_code = ?", code, langCode).First(&translation)
	return &translation, result.Error
}

func (r *translationRepository) GetAllByCode(code string) ([]data.Translation, error) {
	var translation []data.Translation
	result := r.DB.Where("code = ?", code).Find(&translation)
	return translation, result.Error
}

func (r *translationRepository) GetAllDistinctCodesByDomain(domain string) ([]string, error) {
	var codes []string
	result := r.DB.Model(&data.Translation{}).Select("DISTINCT code").Where("domain = ?", domain).Pluck("code", &codes)
	return codes, result.Error
}
