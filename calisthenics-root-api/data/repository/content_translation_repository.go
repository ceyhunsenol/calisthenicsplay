package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IContentTranslationRepository interface {
	Save(translation data.ContentTranslation) (*data.ContentTranslation, error)
	GetAll() ([]data.ContentTranslation, error)
	GetByID(id string) (*data.ContentTranslation, error)
	ExistsByCodeAndLangCode(code, langCode string) (bool, error)
	GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error)
	GetAllByCode(code string) ([]data.ContentTranslation, error)
	Update(translation data.ContentTranslation) (*data.ContentTranslation, error)
	Delete(id string) error
	DeleteAllByContentID(contentID string) error
}

type contentTranslationRepository struct {
	DB *gorm.DB
}

func NewContentTranslationRepository(db *gorm.DB) IContentTranslationRepository {
	return &contentTranslationRepository{DB: db}
}

func (r *contentTranslationRepository) Save(translation data.ContentTranslation) (*data.ContentTranslation, error) {
	result := r.DB.Create(&translation)
	return &translation, result.Error
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

func (r *contentTranslationRepository) ExistsByCodeAndLangCode(code, langCode string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.ContentTranslation{}).Where("code = ? and lang_code = ?", code, langCode).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
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

func (r *contentTranslationRepository) Update(translation data.ContentTranslation) (*data.ContentTranslation, error) {
	result := r.DB.Updates(&translation)
	if result.Error != nil {
		return nil, result.Error
	}
	return &translation, nil
}

func (r *contentTranslationRepository) Delete(id string) error {
	return r.DB.Delete(&data.ContentTranslation{}, "id = ?", id).Error
}

func (r *contentTranslationRepository) DeleteAllByContentID(contentID string) error {
	return r.DB.Delete(&data.ContentTranslation{}, "content_id = ?", contentID).Error
}
