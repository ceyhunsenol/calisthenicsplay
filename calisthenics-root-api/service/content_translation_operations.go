package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"net/http"
)

type IContentTranslationOperations interface {
	SaveContentTranslations(translations []model.ContentTranslationRequest) *model.ServiceError
}

type contentTranslationOperations struct {
	contentTranslationService IContentTranslationService
}

func NewContentTranslationOperations(contentTranslationService IContentTranslationService) IContentTranslationOperations {
	return &contentTranslationOperations{
		contentTranslationService: contentTranslationService,
	}
}

func (o *contentTranslationOperations) SaveContentTranslations(translations []model.ContentTranslationRequest) *model.ServiceError {
	for _, translation := range translations {
		err := o.contentTranslationService.DeleteAllByContentID(translation.ContentID)
		if err != nil {
			return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
		}
		contentTranslation := data.ContentTranslation{
			Code:      translation.Code,
			LangCode:  translation.LangCode,
			Translate: translation.Translate,
			Active:    translation.Active,
			ContentID: translation.ContentID,
		}
		_, err = o.contentTranslationService.Save(contentTranslation)
		if err != nil {
			return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
		}
	}
	return nil
}
