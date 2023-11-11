package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
	"gorm.io/gorm"
)

type IContentService interface {
	Save(tx *gorm.DB, content data.Content) (*data.Content, error)
	GetAll() ([]data.Content, error)
	GetByID(id string) (*data.Content, error)
	ExistsByCode(code string) (bool, error)
	GetByCode(code string) (*data.Content, error)
	Update(tx *gorm.DB, content data.Content) (*data.Content, error)
	Delete(tx *gorm.DB, id string) error
}

type contentService struct {
	contentRepository repository.IContentRepository
}

func NewContentService(contentRepo repository.IContentRepository) IContentService {
	return &contentService{contentRepository: contentRepo}
}

func (s *contentService) Save(tx *gorm.DB, content data.Content) (*data.Content, error) {
	return s.contentRepository.Save(tx, content)
}

func (s *contentService) GetAll() ([]data.Content, error) {
	return s.contentRepository.GetAll()
}

func (s *contentService) GetByID(id string) (*data.Content, error) {
	return s.contentRepository.GetByID(id)
}

func (s *contentService) ExistsByCode(code string) (bool, error) {
	return s.contentRepository.ExistsByCode(code)
}

func (s *contentService) GetByCode(code string) (*data.Content, error) {
	return s.contentRepository.GetByCode(code)
}

func (s *contentService) Update(tx *gorm.DB, content data.Content) (*data.Content, error) {
	return s.contentRepository.Update(tx, content)
}

func (s *contentService) Delete(tx *gorm.DB, id string) error {
	return s.contentRepository.Delete(tx, id)
}
