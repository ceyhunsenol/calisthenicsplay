package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IContentAccessCacheOperations interface {
	SaveCacheContentAccessList() *model.ServiceError
	SaveCacheContentAccess(ID string) (cache.ContentAccessCache, *model.ServiceError)
}

type contentAccessCacheOperations struct {
	contentAccessCacheService cache.IContentAccessCacheService
	contentAccessService      IContentAccessService
}

func NewContentAccessCacheOperations(
	contentAccessCacheService cache.IContentAccessCacheService,
	contentAccessService IContentAccessService,
) IContentAccessCacheOperations {
	return &contentAccessCacheOperations{
		contentAccessCacheService: contentAccessCacheService,
		contentAccessService:      contentAccessService,
	}
}

func (o *contentAccessCacheOperations) SaveCacheContentAccessList() *model.ServiceError {
	contentAccessList, err := o.contentAccessService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	contentAccessCache := make([]cache.ContentAccessCache, 0)
	for _, value := range contentAccessList {
		cac := cache.ContentAccessCache{
			ID:        value.ID,
			ContentID: value.ContentID,
			Audience:  value.Audience,
		}
		contentAccessCache = append(contentAccessCache, cac)
	}
	o.contentAccessCacheService.SaveAllSlice(contentAccessCache)
	return nil
}

func (o *contentAccessCacheOperations) SaveCacheContentAccess(ID string) (cache.ContentAccessCache, *model.ServiceError) {
	access, err := o.contentAccessService.GetByID(ID)
	if err != nil {
		return cache.ContentAccessCache{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Not found"}
	}
	accessCache := cache.ContentAccessCache{
		ContentID: access.ContentID,
		Audience:  access.Audience,
	}
	o.contentAccessCacheService.Save(accessCache)
	return accessCache, nil
}
