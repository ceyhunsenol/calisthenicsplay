package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IRequirementContentService interface {
	Save(requirementContent data.RequirementContent) (*data.RequirementContent, error)
	Delete(requirementContent data.RequirementContent) error
}

type requirementContentService struct {
	requirementContentRepository repository.IRequirementContentRepository
}

func NewRequirementContentService(requirementContentRepository repository.IRequirementContentRepository) IRequirementContentService {
	return &requirementContentService{
		requirementContentRepository: requirementContentRepository,
	}
}

func (s *requirementContentService) Save(requirementContent data.RequirementContent) (*data.RequirementContent, error) {
	return s.requirementContentRepository.Save(requirementContent)
}

func (s *requirementContentService) Delete(requirementContent data.RequirementContent) error {
	return s.requirementContentRepository.Delete(requirementContent)
}
