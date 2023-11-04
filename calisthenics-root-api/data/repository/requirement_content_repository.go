package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IRequirementContentRepository interface {
	Save(requirementContent data.RequirementContent) (*data.RequirementContent, error)
	Delete(requirementContent data.RequirementContent) error
}

type requirementContentRepository struct {
	DB *gorm.DB
}

func NewRequirementContentRepository(db *gorm.DB) IRequirementContentRepository {
	return &requirementContentRepository{DB: db}
}

func (r *requirementContentRepository) Save(requirementContent data.RequirementContent) (*data.RequirementContent, error) {
	result := r.DB.Create(&requirementContent)
	return &requirementContent, result.Error
}

func (r *requirementContentRepository) Delete(requirementContent data.RequirementContent) error {
	return r.DB.Delete(&data.RequirementContent{}, "content_id = ? AND requirement_content_id = ?", requirementContent.ContentID, requirementContent.RequirementContentID).Error
}
