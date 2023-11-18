package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type ITranslationService interface {
	GetAll() ([]data.Translation, error)
	GetByID(id string) (*data.Translation, error)
	GetByCodeAndLangCode(code, langCode string) (*data.Translation, error)
	GetAllByCode(code string) ([]data.Translation, error)
	GetAllDistinctCodesByDomain(domain string) ([]string, error)
}

type TranslationService struct {
	translationRepository repository.ITranslationRepository
}

func NewTranslationService(repo repository.ITranslationRepository) ITranslationService {
	return &TranslationService{
		translationRepository: repo,
	}
}

func (s *TranslationService) GetAll() ([]data.Translation, error) {
	return s.translationRepository.GetAll()
}

func (s *TranslationService) GetByID(id string) (*data.Translation, error) {
	return s.translationRepository.GetByID(id)
}

func (s *TranslationService) GetByCodeAndLangCode(code, langCode string) (*data.Translation, error) {
	return s.translationRepository.GetByCodeAndLangCode(code, langCode)
}

func (s *TranslationService) GetAllByCode(code string) ([]data.Translation, error) {
	return s.translationRepository.GetAllByCode(code)
}

func (s *TranslationService) GetAllDistinctCodesByDomain(domain string) ([]string, error) {
	return s.translationRepository.GetAllDistinctCodesByDomain(domain)
}
