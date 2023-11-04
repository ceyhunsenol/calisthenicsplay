package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IMediaService interface {
	GetAll() ([]data.Media, error)
	GetByID(id string) (*data.Media, error)
}

type mediaService struct {
	mediaRepository repository.IMediaRepository
}

func NewMediaService(mediaRepo repository.IMediaRepository) IMediaService {
	return &mediaService{mediaRepository: mediaRepo}
}

func (s *mediaService) GetAll() ([]data.Media, error) {
	return s.mediaRepository.GetAll()
}

func (s *mediaService) GetByID(id string) (*data.Media, error) {
	return s.mediaRepository.GetByID(id)
}
