package repository

import (
	"calisthenics-content-api/data"
	"gorm.io/gorm"
)

type IContentAccessRepository interface {
	Save(contentAccess data.ContentAccess) (*data.ContentAccess, error)
	GetAll() ([]data.ContentAccess, error)
	GetByID(id string) (*data.ContentAccess, error)
	ExistsByContentID(contentID string) (bool, error)
	GetByContentID(contentID string) (*data.ContentAccess, error)
	Update(contentAccess data.ContentAccess) (*data.ContentAccess, error)
	Delete(id string) error
}

type contentAccessRepository struct {
	DB *gorm.DB
}

func NewContentAccessRepository(db *gorm.DB) IContentAccessRepository {
	return &contentAccessRepository{DB: db}
}

func (r *contentAccessRepository) Save(contentAccess data.ContentAccess) (*data.ContentAccess, error) {
	result := r.DB.Create(&contentAccess)
	return &contentAccess, result.Error
}

func (r *contentAccessRepository) GetAll() ([]data.ContentAccess, error) {
	var contentAccessList []data.ContentAccess
	result := r.DB.Find(&contentAccessList)
	return contentAccessList, result.Error
}

func (r *contentAccessRepository) GetByID(id string) (*data.ContentAccess, error) {
	var contentAccess data.ContentAccess
	result := r.DB.Where("id = ?", id).First(&contentAccess)
	return &contentAccess, result.Error
}

func (r *contentAccessRepository) ExistsByContentID(contentID string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.ContentAccess{}).Where("content_id = ?", contentID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *contentAccessRepository) GetByContentID(contentID string) (*data.ContentAccess, error) {
	var contentAccess data.ContentAccess
	result := r.DB.Where("content_id = ?", contentID).First(&contentAccess)
	return &contentAccess, result.Error
}

func (r *contentAccessRepository) Update(contentAccess data.ContentAccess) (*data.ContentAccess, error) {
	result := r.DB.Updates(&contentAccess)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contentAccess, nil
}

func (r *contentAccessRepository) Delete(id string) error {
	return r.DB.Delete(&data.ContentAccess{}, "id = ?", id).Error
}
