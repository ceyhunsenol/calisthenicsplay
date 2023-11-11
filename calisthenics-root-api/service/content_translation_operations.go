package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"gorm.io/gorm"
	"net/http"
)

type IContentTranslationOperations interface {
	SaveContentTranslations(tx *gorm.DB, request model.ContentTranslationRequest) *model.ServiceError
	DeleteAllContentTranslations(tx *gorm.DB, contentID string) *model.ServiceError
}

type contentTranslationOperations struct {
	contentTranslationService IContentTranslationService
}

func NewContentTranslationOperations(contentTranslationService IContentTranslationService) IContentTranslationOperations {
	return &contentTranslationOperations{
		contentTranslationService: contentTranslationService,
	}
}

func (o *contentTranslationOperations) SaveContentTranslations(tx *gorm.DB, request model.ContentTranslationRequest) *model.ServiceError {
	err := o.contentTranslationService.DeleteAllByContentID(tx, request.ContentID)
	if err != nil {
		tx.Rollback()
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
		_, err = o.contentTranslationService.Save(tx, contentTranslation)
		if err != nil {
			tx.Rollback()
			return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
		}
	}
	return nil
}

func (o *contentTranslationOperations) DeleteAllContentTranslations(tx *gorm.DB, contentID string) *model.ServiceError {
	err := o.contentTranslationService.DeleteAllByContentID(tx, contentID)
	if err != nil {
		tx.Rollback()
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "failed to translation process."}
	}
	return nil
}
