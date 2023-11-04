package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"net/http"
)

type IHelperContentOperations interface {
	AddHelperContent(request model.HelperContentRequest) *model.ServiceError
	RemoveHelperContent(request model.HelperContentRequest) *model.ServiceError
}

type helperContentOperations struct {
	helperContentService IHelperContentService
	contentService       IContentService
}

func NewHelperContentOperations(helperContentService IHelperContentService, contentService IContentService) IHelperContentOperations {
	return &helperContentOperations{
		helperContentService: helperContentService,
		contentService:       contentService,
	}
}

func (u *helperContentOperations) AddHelperContent(request model.HelperContentRequest) *model.ServiceError {
	content, err := u.contentService.GetByID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found."}
	}

	requirementContent, err := u.contentService.GetByID(request.HelperContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Helper content not found."}
	}

	_, err = u.helperContentService.Save(data.HelperContent{
		ContentID:       content.ID,
		HelperContentID: requirementContent.ID,
	})

	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Content could not be added."}
	}

	return nil
}

func (u *helperContentOperations) RemoveHelperContent(request model.HelperContentRequest) *model.ServiceError {
	content, err := u.contentService.GetByID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found."}
	}

	requirementContent, err := u.contentService.GetByID(request.HelperContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Helper content not found."}
	}

	err = u.helperContentService.Delete(data.HelperContent{
		ContentID:       content.ID,
		HelperContentID: requirementContent.ID,
	})
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Content could not be deleted."}
	}
	return nil
}
