package service

import (
	"calisthenics-auth-api/data"
	"calisthenics-auth-api/data/repository"
)

type IProfileService interface {
	GetByUserID(userID string) (*data.Profile, error)
	SaveOrUpdate(user data.Profile) (*data.Profile, error)
	GetById(id string) (*data.Profile, error)
}

type profileService struct {
	profileRepository repository.IProfileRepository
}

func NewProfileService(profileRepository repository.IProfileRepository) IProfileService {
	return &profileService{profileRepository: profileRepository}
}

func (s *profileService) GetByUserID(userID string) (*data.Profile, error) {
	return s.profileRepository.GetByUserID(userID)
}

func (s *profileService) GetById(id string) (*data.Profile, error) {
	return s.profileRepository.GetById(id)
}

func (s *profileService) SaveOrUpdate(profile data.Profile) (*data.Profile, error) {
	return s.profileRepository.SaveOrUpdate(profile)
}
