package repository

import (
	"calisthenics-auth-api/data"
	"gorm.io/gorm"
)

type IProfileRepository interface {
	GetById(id string) (*data.Profile, error)
	SaveOrUpdate(profile data.Profile) (*data.Profile, error)
	GetByUserID(userID string) (*data.Profile, error)
}

type profileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) IProfileRepository {
	return &profileRepository{DB: db}
}

func (r *profileRepository) GetById(id string) (*data.Profile, error) {
	var profile data.Profile
	result := r.DB.First(&profile, id)
	return &profile, result.Error
}

func (r *profileRepository) SaveOrUpdate(profile data.Profile) (*data.Profile, error) {
	var existingProfile data.Profile
	if err := r.DB.Where("user_id = ?", profile.UserID).First(&existingProfile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := r.DB.Create(&profile).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		profile.ID = existingProfile.ID
		if err := r.DB.Save(&profile).Error; err != nil {
			return nil, err
		}
	}
	return &profile, nil
}

func (r *profileRepository) GetByUserID(userID string) (*data.Profile, error) {
	var profile data.Profile
	result := r.DB.Where("user_id = ?", userID).First(&profile)
	return &profile, result.Error
}
