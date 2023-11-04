package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IHelperContentRepository interface {
	Save(content data.HelperContent) (*data.HelperContent, error)
	Delete(helperContent data.HelperContent) error
}

type helperContentRepository struct {
	DB *gorm.DB
}

func NewHelperContentRepository(db *gorm.DB) IHelperContentRepository {
	return &helperContentRepository{DB: db}
}

func (r *helperContentRepository) Save(helperContent data.HelperContent) (*data.HelperContent, error) {
	result := r.DB.Create(&helperContent)
	return &helperContent, result.Error
}

func (r *helperContentRepository) Delete(helperContent data.HelperContent) error {
	return r.DB.Delete(&data.HelperContent{}, "content_id = ? AND helper_content_id = ?", helperContent.ContentID, helperContent.HelperContentID).Error
}
