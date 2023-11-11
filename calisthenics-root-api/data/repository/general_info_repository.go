package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IGeneralInfoRepository interface {
	GetByID(id string) (*data.GeneralInfo, error)
}

type generalInfoRepository struct {
	DB *gorm.DB
}

func NewGeneralInfoRepository(db *gorm.DB) IGeneralInfoRepository {
	return &generalInfoRepository{DB: db}
}

func (r *generalInfoRepository) GetAll() ([]data.GeneralInfo, error) {
	var contents []data.GeneralInfo
	result := r.DB.Preload("Medias").Find(&contents)
	return contents, result.Error
}

func (r *generalInfoRepository) GetByID(id string) (*data.GeneralInfo, error) {
	var content data.GeneralInfo
	result := r.DB.Where("id = ?", id).First(&content)
	return &content, result.Error
}

func (r *generalInfoRepository) GetByKey(key string) (*data.GeneralInfo, error) {
	var content data.GeneralInfo
	result := r.DB.Where("info_key = ?", key).First(&content)
	return &content, result.Error
}
