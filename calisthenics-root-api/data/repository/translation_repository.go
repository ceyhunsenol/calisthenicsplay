package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type ITranslationRepository interface {
	Save(translation data.Translation) (*data.Translation, error)
	GetAll() ([]data.Translation, error)
	GetByID(id string) (*data.Translation, error)
	ExistsByCodeAndLangCode(code, langCode string) (bool, error)
	GetByCodeAndLangCode(id, langCode string) (*data.Translation, error)
	GetAllByCode(code string) ([]data.Translation, error)
	Update(translation data.Translation) (*data.Translation, error)
	Delete(id string) error
}

type translationRepository struct {
	DB *gorm.DB
}

func NewTranslationRepository(db *gorm.DB) ITranslationRepository {
	return &translationRepository{DB: db}
}

func (r *translationRepository) Save(translation data.Translation) (*data.Translation, error) {
	result := r.DB.Create(&translation)
	return &translation, result.Error
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

func (r *translationRepository) ExistsByCodeAndLangCode(code, langCode string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.Translation{}).Where("code = ? and lang_code = ?", code, langCode).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
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

func (r *translationRepository) Update(translation data.Translation) (*data.Translation, error) {
	result := r.DB.Updates(&translation)
	if result.Error != nil {
		return nil, result.Error
	}
	return &translation, nil
}

func (r *translationRepository) Delete(id string) error {
	return r.DB.Delete(&data.Translation{}, "id = ?", id).Error
}
