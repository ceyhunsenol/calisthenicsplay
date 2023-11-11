package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
	"gorm.io/gorm"
)

type IMediaService interface {
	Save(tx *gorm.DB, media data.Media) (*data.Media, error)
	GetAll() ([]data.Media, error)
	GetByID(id string) (*data.Media, error)
	Update(tx *gorm.DB, media data.Media) (*data.Media, error)
	Delete(tx *gorm.DB, id string) error
}

type mediaService struct {
	mediaRepository repository.IMediaRepository
}

func NewMediaService(mediaRepo repository.IMediaRepository) IMediaService {
	return &mediaService{mediaRepository: mediaRepo}
}

func (s *mediaService) Save(tx *gorm.DB, media data.Media) (*data.Media, error) {
	return s.mediaRepository.Save(tx, media)
}

func (s *mediaService) GetAll() ([]data.Media, error) {
	return s.mediaRepository.GetAll()
}

func (s *mediaService) GetByID(id string) (*data.Media, error) {
	return s.mediaRepository.GetByID(id)
}

func (s *mediaService) Update(tx *gorm.DB, media data.Media) (*data.Media, error) {
	return s.mediaRepository.Update(tx, media)
}

func (s *mediaService) Delete(tx *gorm.DB, id string) error {
	return s.mediaRepository.Delete(tx, id)
}
