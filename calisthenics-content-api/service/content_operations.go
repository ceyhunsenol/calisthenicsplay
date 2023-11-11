package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IContentOperations interface {
	GetContentByCode(code string) (model.ContentModel, *model.ServiceError)
}

type contentOperations struct {
	contentService      IContentService
	contentCacheService cache.IContentCacheService
}

func NewContentOperations(contentService IContentService, contentCacheService cache.IContentCacheService) IContentOperations {
	return &contentOperations{
		contentService:      contentService,
		contentCacheService: contentCacheService,
	}
}

func (o *contentOperations) GetContentByCode(code string) (model.ContentModel, *model.ServiceError) {
	content, err := o.contentCacheService.GetByCode(code)
	if err != nil {
		return model.ContentModel{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found"}
	}
	contentModel := model.ContentModel{ID: content.ID}
	return contentModel, nil
}
