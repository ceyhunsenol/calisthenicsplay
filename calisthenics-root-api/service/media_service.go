package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IMediaService interface {
	Save(media data.Media) (*data.Media, error)
	GetAll() ([]data.Media, error)
	GetByID(id string) (*data.Media, error)
	Update(media data.Media) (*data.Media, error)
	Delete(id string) error
}

type mediaService struct {
	mediaRepository repository.IMediaRepository
}

func NewMediaService(mediaRepo repository.IMediaRepository) IMediaService {
	return &mediaService{mediaRepository: mediaRepo}
}

func (s *mediaService) Save(media data.Media) (*data.Media, error) {
	return s.mediaRepository.Save(media)
}

func (s *mediaService) GetAll() ([]data.Media, error) {
	return s.mediaRepository.GetAll()
}

func (s *mediaService) GetByID(id string) (*data.Media, error) {
	return s.mediaRepository.GetByID(id)
}

func (s *mediaService) Update(media data.Media) (*data.Media, error) {
	return s.mediaRepository.Update(media)
}

func (s *mediaService) Delete(id string) error {
	return s.mediaRepository.Delete(id)
}
