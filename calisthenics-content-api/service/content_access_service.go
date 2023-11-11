package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IContentAccessService interface {
	Save(contentAccess data.ContentAccess) (*data.ContentAccess, error)
	GetAll() ([]data.ContentAccess, error)
	GetByID(id string) (*data.ContentAccess, error)
	ExistsByContentID(contentID string) (bool, error)
	GetByContentID(contentID string) (*data.ContentAccess, error)
	Update(contentAccess data.ContentAccess) (*data.ContentAccess, error)
	Delete(id string) error
}

type contentAccessService struct {
	contentAccessRepository repository.IContentAccessRepository
}

func NewContentAccessService(contentAccessRepository repository.IContentAccessRepository) IContentAccessService {
	return &contentAccessService{contentAccessRepository: contentAccessRepository}
}

func (s *contentAccessService) Save(contentAccess data.ContentAccess) (*data.ContentAccess, error) {
	return s.contentAccessRepository.Save(contentAccess)
}

func (s *contentAccessService) GetAll() ([]data.ContentAccess, error) {
	return s.contentAccessRepository.GetAll()
}

func (s *contentAccessService) GetByID(id string) (*data.ContentAccess, error) {
	return s.contentAccessRepository.GetByID(id)
}

func (s *contentAccessService) ExistsByContentID(contentID string) (bool, error) {
	return s.contentAccessRepository.ExistsByContentID(contentID)
}

func (s *contentAccessService) GetByContentID(contentID string) (*data.ContentAccess, error) {
	return s.contentAccessRepository.GetByContentID(contentID)
}

func (s *contentAccessService) Update(contentAccess data.ContentAccess) (*data.ContentAccess, error) {
	return s.contentAccessRepository.Update(contentAccess)
}

func (s *contentAccessService) Delete(id string) error {
	return s.contentAccessRepository.Delete(id)
}
