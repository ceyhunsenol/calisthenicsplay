package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type ITranslationService interface {
	Save(translation data.Translation) (*data.Translation, error)
	GetAll() ([]data.Translation, error)
	GetByID(id string) (*data.Translation, error)
	ExistsByCodeAndLangCode(code, langCode string) (bool, error)
	GetByCodeAndLangCode(code, langCode string) (*data.Translation, error)
	GetAllByCode(code string) ([]data.Translation, error)
	Update(translation data.Translation) (*data.Translation, error)
	Delete(id string) error
}

type TranslationService struct {
	translationRepository repository.ITranslationRepository
}

func NewTranslationService(repo repository.ITranslationRepository) ITranslationService {
	return &TranslationService{
		translationRepository: repo,
	}
}

func (s *TranslationService) Save(translation data.Translation) (*data.Translation, error) {
	return s.translationRepository.Save(translation)
}

func (s *TranslationService) GetAll() ([]data.Translation, error) {
	return s.translationRepository.GetAll()
}

func (s *TranslationService) GetByID(id string) (*data.Translation, error) {
	return s.translationRepository.GetByID(id)
}

func (s *TranslationService) ExistsByCodeAndLangCode(code, langCode string) (bool, error) {
	return s.translationRepository.ExistsByCodeAndLangCode(code, langCode)
}

func (s *TranslationService) GetByCodeAndLangCode(code, langCode string) (*data.Translation, error) {
	return s.translationRepository.GetByCodeAndLangCode(code, langCode)
}

func (s *TranslationService) GetAllByCode(code string) ([]data.Translation, error) {
	return s.translationRepository.GetAllByCode(code)
}

func (s *TranslationService) Update(translation data.Translation) (*data.Translation, error) {
	return s.translationRepository.Update(translation)
}

func (s *TranslationService) Delete(id string) error {
	return s.translationRepository.Delete(id)
}
