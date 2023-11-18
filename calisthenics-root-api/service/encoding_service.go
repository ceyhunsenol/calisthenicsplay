package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IEncodingService interface {
	Save(encoding data.Encoding) (*data.Encoding, error)
	GetAll() ([]data.Encoding, error)
	GetByID(id string) (*data.Encoding, error)
	Update(encoding data.Encoding) (*data.Encoding, error)
	Delete(id string) error
}

type encodingService struct {
	encodingRepository repository.IEncodingRepository
}

func NewEncodingService(encodingRepository repository.IEncodingRepository) IEncodingService {
	return &encodingService{encodingRepository: encodingRepository}
}

func (s *encodingService) Save(encoding data.Encoding) (*data.Encoding, error) {
	return s.encodingRepository.Save(encoding)
}

func (s *encodingService) GetAll() ([]data.Encoding, error) {
	return s.encodingRepository.GetAll()
}

func (s *encodingService) GetByID(id string) (*data.Encoding, error) {
	return s.encodingRepository.GetByID(id)
}

func (s *encodingService) Update(encoding data.Encoding) (*data.Encoding, error) {
	return s.encodingRepository.Update(encoding)
}

func (s *encodingService) Delete(id string) error {
	return s.encodingRepository.Delete(id)
}
