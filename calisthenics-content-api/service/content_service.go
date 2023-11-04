package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IContentService interface {
	GetAll() ([]data.Content, error)
	GetByID(id string) (*data.Content, error)
	GetByCode(code string) (*data.Content, error)
}

type contentService struct {
	contentRepository repository.IContentRepository
}

func NewContentService(contentRepo repository.IContentRepository) IContentService {
	return &contentService{contentRepository: contentRepo}
}

func (s *contentService) GetAll() ([]data.Content, error) {
	return s.contentRepository.GetAll()
}

func (s *contentService) GetByID(id string) (*data.Content, error) {
	return s.contentRepository.GetByID(id)
}

func (s *contentService) GetByCode(code string) (*data.Content, error) {
	return s.contentRepository.GetByCode(code)
}
