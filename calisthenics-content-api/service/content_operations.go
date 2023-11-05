package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IContentOperations interface {
	SaveCacheContents() error
	SaveCacheContent(ID string) (cache.ContentCache, error)
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

func (o *contentOperations) SaveCacheContents() error {
	contents, err := o.contentService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "General error"}
	}

	activeContents := make([]cache.ContentCache, 0)
	for _, value := range contents {
		if value.Active {
			contentCache := cache.ContentCache{
				ID:                    value.ID,
				CodeMultiLang:         nil,
				DescriptionMultiLang:  nil,
				Active:                true,
				HelperContentIDs:      nil,
				RequirementContentIDs: nil,
			}
			activeContents = append(activeContents, contentCache)
		}
	}
	o.contentCacheService.SaveAllSlice(activeContents)
	return nil
}

func (o *contentOperations) SaveCacheContent(ID string) (cache.ContentCache, error) {
	content, err := o.contentService.GetByID(ID)
	if err != nil {
		return cache.ContentCache{}, &model.ServiceError{Code: http.StatusInternalServerError, Message: "General error"}
	}
	contentCache := cache.ContentCache{
		ID:                    content.ID,
		CodeMultiLang:         nil,
		DescriptionMultiLang:  nil,
		Active:                true,
		HelperContentIDs:      nil,
		RequirementContentIDs: nil,
	}
	o.contentCacheService.Save(contentCache)
	return contentCache, nil
}
