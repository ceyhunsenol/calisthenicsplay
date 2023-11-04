package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"net/http"
)

type IRequirementContentOperations interface {
	AddRequirementContent(request model.RequirementContentRequest) *model.ServiceError
	RemoveRequirementContent(request model.RequirementContentRequest) *model.ServiceError
}

type requirementContentOperations struct {
	requirementContentService IRequirementContentService
	contentService            IContentService
}

func NewRequirementContentOperations(requirementContentService IRequirementContentService, contentService IContentService) IRequirementContentOperations {
	return &requirementContentOperations{
		requirementContentService: requirementContentService,
		contentService:            contentService,
	}
}

func (u *requirementContentOperations) AddRequirementContent(request model.RequirementContentRequest) *model.ServiceError {
	content, err := u.contentService.GetByID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found."}
	}

	requirementContent, err := u.contentService.GetByID(request.RequirementContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Requirement content not found."}
	}

	_, err = u.requirementContentService.Save(data.RequirementContent{
		ContentID:            content.ID,
		RequirementContentID: requirementContent.ID,
	})

	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Content could not be added."}
	}

	return nil
}

func (u *requirementContentOperations) RemoveRequirementContent(request model.RequirementContentRequest) *model.ServiceError {
	content, err := u.contentService.GetByID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found."}
	}

	requirementContent, err := u.contentService.GetByID(request.RequirementContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Requirement content not found."}
	}

	err = u.requirementContentService.Delete(data.RequirementContent{
		ContentID:            content.ID,
		RequirementContentID: requirementContent.ID,
	})
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Content could not be deleted."}
	}
	return nil
}
