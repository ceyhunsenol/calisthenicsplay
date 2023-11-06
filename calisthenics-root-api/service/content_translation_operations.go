package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"net/http"
)

type IContentTranslationOperations interface {
	SaveContentTranslations(request model.ContentTranslationRequest) *model.ServiceError
	DeleteAllContentTranslations(contentID string) *model.ServiceError
}

type contentTranslationOperations struct {
	contentTranslationService IContentTranslationService
}

func NewContentTranslationOperations(contentTranslationService IContentTranslationService) IContentTranslationOperations {
	return &contentTranslationOperations{
		contentTranslationService: contentTranslationService,
	}
}

func (o *contentTranslationOperations) SaveContentTranslations(request model.ContentTranslationRequest) *model.ServiceError {
	err := o.contentTranslationService.DeleteAllByContentID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
	}
	for _, translation := range request.Translations {
		contentTranslation := data.ContentTranslation{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			ContentID: request.ContentID,
		}
		_, err = o.contentTranslationService.Save(contentTranslation)
		if err != nil {
			return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
		}
	}
	return nil
}

func (o *contentTranslationOperations) DeleteAllContentTranslations(contentID string) *model.ServiceError {
	err := o.contentTranslationService.DeleteAllByContentID(contentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
	}
	return nil
}
