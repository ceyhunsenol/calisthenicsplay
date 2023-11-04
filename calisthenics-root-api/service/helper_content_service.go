package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IHelperContentService interface {
	Save(helperContent data.HelperContent) (*data.HelperContent, error)
	Delete(helperContent data.HelperContent) error
}

type helperContentService struct {
	contentService          IContentService
	helperContentRepository repository.IHelperContentRepository
}

func NewHelperContentService(helperContentRepository repository.IHelperContentRepository, contentService IContentService) IHelperContentService {
	return &helperContentService{
		contentService:          contentService,
		helperContentRepository: helperContentRepository,
	}
}

func (u *helperContentService) Save(helperContent data.HelperContent) (*data.HelperContent, error) {
	return u.helperContentRepository.Save(helperContent)
}

func (u *helperContentService) Delete(helperContent data.HelperContent) error {
	return u.helperContentRepository.Delete(helperContent)
}
