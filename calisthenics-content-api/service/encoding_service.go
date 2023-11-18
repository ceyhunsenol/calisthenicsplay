package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IEncodingService interface {
	GetAll() ([]data.Encoding, error)
	GetByID(id string) (*data.Encoding, error)
}

type encodingService struct {
	encodingRepository repository.IEncodingRepository
}

func NewEncodingService(encodingRepository repository.IEncodingRepository) IEncodingService {
	return &encodingService{encodingRepository: encodingRepository}
}

func (s *encodingService) GetAll() ([]data.Encoding, error) {
	return s.encodingRepository.GetAll()
}

func (s *encodingService) GetByID(id string) (*data.Encoding, error) {
	return s.encodingRepository.GetByID(id)
}
