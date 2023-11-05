package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IContentTranslationService interface {
	Save(translation data.ContentTranslation) (*data.ContentTranslation, error)
	GetAll() ([]data.ContentTranslation, error)
	GetByID(id string) (*data.ContentTranslation, error)
	ExistsByCodeAndLangCode(code, langCode string) (bool, error)
	GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error)
	GetAllByCode(code string) ([]data.ContentTranslation, error)
	Update(translation data.ContentTranslation) (*data.ContentTranslation, error)
	Delete(id string) error
}

type ContentTranslationService struct {
	contentTranslationRepository repository.IContentTranslationRepository
}

func NewContentTranslationService(repo repository.IContentTranslationRepository) IContentTranslationService {
	return &ContentTranslationService{
		contentTranslationRepository: repo,
	}
}

func (s *ContentTranslationService) Save(translation data.ContentTranslation) (*data.ContentTranslation, error) {
	return s.contentTranslationRepository.Save(translation)
}

func (s *ContentTranslationService) GetAll() ([]data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetAll()
}

func (s *ContentTranslationService) GetByID(id string) (*data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetByID(id)
}

func (s *ContentTranslationService) ExistsByCodeAndLangCode(code, langCode string) (bool, error) {
	return s.contentTranslationRepository.ExistsByCodeAndLangCode(code, langCode)
}

func (s *ContentTranslationService) GetByCodeAndLangCode(code, langCode string) (*data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetByCodeAndLangCode(code, langCode)
}

func (s *ContentTranslationService) GetAllByCode(code string) ([]data.ContentTranslation, error) {
	return s.contentTranslationRepository.GetAllByCode(code)
}

func (s *ContentTranslationService) Update(translation data.ContentTranslation) (*data.ContentTranslation, error) {
	return s.contentTranslationRepository.Update(translation)
}

func (s *ContentTranslationService) Delete(id string) error {
	return s.contentTranslationRepository.Delete(id)
}
