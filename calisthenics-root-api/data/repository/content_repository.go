package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IContentRepository interface {
	Save(tx *gorm.DB, content data.Content) (*data.Content, error)
	GetAll() ([]data.Content, error)
	GetByID(id string) (*data.Content, error)
	GetByCode(code string) (*data.Content, error)
	ExistsByCode(code string) (bool, error)
	Update(tx *gorm.DB, content data.Content) (*data.Content, error)
	Delete(tx *gorm.DB, id string) error
}

type contentRepository struct {
	DB *gorm.DB
}

func NewContentRepository(db *gorm.DB) IContentRepository {
	return &contentRepository{DB: db}
}

func (r *contentRepository) Save(tx *gorm.DB, content data.Content) (*data.Content, error) {
	result := tx.Create(&content)
	return &content, result.Error
}

func (r *contentRepository) GetAll() ([]data.Content, error) {
	var contents []data.Content
	result := r.DB.Preload("Medias").Find(&contents)
	return contents, result.Error
}

func (r *contentRepository) GetByID(id string) (*data.Content, error) {
	var content data.Content
	result := r.DB.Where("id = ?", id).Preload("Medias").Preload("HelperContents").Preload("RequirementContents").First(&content)
	return &content, result.Error
}

func (r *contentRepository) GetByCode(code string) (*data.Content, error) {
	var content data.Content
	result := r.DB.Where("code = ?", code).Preload("Medias").First(&content)
	return &content, result.Error
}

func (r *contentRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.Content{}).Where("code = ?", code).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *contentRepository) Update(tx *gorm.DB, content data.Content) (*data.Content, error) {
	result := tx.Updates(&content)
	if result.Error != nil {
		return nil, result.Error
	}
	return &content, nil
}

func (r *contentRepository) Delete(tx *gorm.DB, id string) error {
	return tx.Delete(&data.Content{}, "id = ?", id).Error
}
