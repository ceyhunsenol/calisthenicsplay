package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IMediaAccessService interface {
	Save(mediaAccess data.MediaAccess) (*data.MediaAccess, error)
	GetAll() ([]data.MediaAccess, error)
	GetByID(id string) (*data.MediaAccess, error)
	ExistsByMediaID(mediaID string) (bool, error)
	GetByMediaID(mediaID string) (*data.MediaAccess, error)
	Update(mediaAccess data.MediaAccess) (*data.MediaAccess, error)
	Delete(id string) error
}

type mediaAccessService struct {
	mediaAccessRepository repository.IMediaAccessRepository
}

func NewMediaAccessService(mediaAccessRepository repository.IMediaAccessRepository) IMediaAccessService {
	return &mediaAccessService{mediaAccessRepository: mediaAccessRepository}
}

func (s *mediaAccessService) Save(mediaAccess data.MediaAccess) (*data.MediaAccess, error) {
	return s.mediaAccessRepository.Save(mediaAccess)
}

func (s *mediaAccessService) GetAll() ([]data.MediaAccess, error) {
	return s.mediaAccessRepository.GetAll()
}

func (s *mediaAccessService) GetByID(id string) (*data.MediaAccess, error) {
	return s.mediaAccessRepository.GetByID(id)
}

func (s *mediaAccessService) ExistsByMediaID(mediaID string) (bool, error) {
	return s.mediaAccessRepository.ExistsByMediaID(mediaID)
}

func (s *mediaAccessService) GetByMediaID(mediaID string) (*data.MediaAccess, error) {
	return s.mediaAccessRepository.GetByMediaID(mediaID)
}

func (s *mediaAccessService) Update(mediaAccess data.MediaAccess) (*data.MediaAccess, error) {
	return s.mediaAccessRepository.Update(mediaAccess)
}

func (s *mediaAccessService) Delete(id string) error {
	return s.mediaAccessRepository.Delete(id)
}
