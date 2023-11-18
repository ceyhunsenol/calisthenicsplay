package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IContentTranslationService interface {
	GetAll() ([]data.ContentTranslation, error)
	GetByID(id string) (*data.ContentTranslation, error)
	GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error)
	GetAllByCode(code string) ([]data.ContentTranslation, error)
}

type ContentTranslationService struct {
	contentTranslationRepository repository.IContentTranslationRepository
}

func NewContentTranslationService(repo repository.IContentTranslationRepository) IContentTranslationService {
	return &ContentTranslationService{
		contentTranslationRepository: repo,
	}
}

func (s *ContentTranslationService) GetAll() ([]data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetAll()
}

func (s *ContentTranslationService) GetByID(id string) (*data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetByID(id)
}

func (s *ContentTranslationService) GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetByCodeAndLangCode(code, langCode)
}

func (s *ContentTranslationService) GetAllByCode(code string) ([]data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetAllByCode(code)
}
