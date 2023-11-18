package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IMediaAccessCacheOperations interface {
	SaveCacheMediaAccessList() *model.ServiceError
	SaveCacheMediaAccess(ID string) interface{}
}

type mediaAccessCacheOperations struct {
	mediaAccessCacheService cache.IMediaAccessCacheService
	mediaAccessService      IMediaAccessService
}

func NewMediaAccessCacheOperations(
	mediaAccessCacheService cache.IMediaAccessCacheService,
	mediaAccessService IMediaAccessService,
) IMediaAccessCacheOperations {
	return &mediaAccessCacheOperations{
		mediaAccessCacheService: mediaAccessCacheService,
		mediaAccessService:      mediaAccessService,
	}
}

func (o *mediaAccessCacheOperations) SaveCacheMediaAccessList() *model.ServiceError {
	o.mediaAccessCacheService.RemoveAll()
	mediaAccessList, err := o.mediaAccessService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	mediaAccessCache := make([]cache.MediaAccessCache, 0)
	for _, value := range mediaAccessList {
		cac := cache.MediaAccessCache{
			MediaID:  value.MediaID,
			Audience: value.Audience,
		}
		mediaAccessCache = append(mediaAccessCache, cac)
	}
	o.mediaAccessCacheService.SaveAllSlice(mediaAccessCache)
	return nil
}

func (o *mediaAccessCacheOperations) SaveCacheMediaAccess(ID string) interface{} {
	o.mediaAccessCacheService.Remove(ID)
	access, err := o.mediaAccessService.GetByID(ID)
	if err != nil {
		return nil
	}
	accessCache := cache.MediaAccessCache{
		MediaID:  access.MediaID,
		Audience: access.Audience,
	}
	o.mediaAccessCacheService.Save(accessCache)
	return &accessCache
}
