package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IEncodingFileService interface {
	Save(encodingFile data.EncodingFile) (*data.EncodingFile, error)
	GetAllEncodingFilesByEncodingID(encodingID string) ([]data.EncodingFile, error)
	GetByID(id string) (*data.EncodingFile, error)
	Update(encodingFile data.EncodingFile) (*data.EncodingFile, error)
	Delete(id string) error
	DeleteEncodingFilesByEncodingID(encodingID string) error
}

type encodingFileService struct {
	encodingFileRepository repository.IEncodingFileRepository
}

func NewEncodingFileService(encodingFileRepository repository.IEncodingFileRepository) IEncodingFileService {
	return &encodingFileService{encodingFileRepository: encodingFileRepository}
}

func (s *encodingFileService) Save(encodingFile data.EncodingFile) (*data.EncodingFile, error) {
	return s.encodingFileRepository.Save(encodingFile)
}

func (s *encodingFileService) GetAllEncodingFilesByEncodingID(encodingID string) ([]data.EncodingFile, error) {
	return s.encodingFileRepository.GetAllEncodingFilesByEncodingID(encodingID)
}

func (s *encodingFileService) GetByID(id string) (*data.EncodingFile, error) {
	return s.encodingFileRepository.GetByID(id)
}

func (s *encodingFileService) Update(encodingFile data.EncodingFile) (*data.EncodingFile, error) {
	return s.encodingFileRepository.Update(encodingFile)
}

func (s *encodingFileService) Delete(id string) error {
	return s.encodingFileRepository.Delete(id)
}

func (s *encodingFileService) DeleteEncodingFilesByEncodingID(id string) error {
	return s.encodingFileRepository.DeleteEncodingFilesByEncodingID(id)
}
