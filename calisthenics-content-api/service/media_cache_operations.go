package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/data"
	"calisthenics-content-api/model"
	"calisthenics-content-api/pkg"
	"net/http"
)

type IMediaCacheOperations interface {
	SaveCacheMedias() *model.ServiceError
	SaveCacheMedia(ID string) (cache.MediaCache, *model.ServiceError)
}

type mediaCacheOperations struct {
	mediaService      IMediaService
	mediaCacheService cache.IMediaCacheService
	contentService    IContentService
}

func NewMediaCacheOperations(mediaService IMediaService, mediaCacheService cache.IMediaCacheService, contentService IContentService) IMediaCacheOperations {
	return &mediaCacheOperations{
		mediaService:      mediaService,
		mediaCacheService: mediaCacheService,
		contentService:    contentService,
	}
}

func (o *mediaCacheOperations) SaveCacheMedias() *model.ServiceError {
	medias, err := o.mediaService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	grouped := pkg.GroupByField(medias, func(media *data.Media) string {
		return media.ContentID
	})

	for key, value := range grouped {
		content, contentError := o.contentService.GetByID(key)
		if contentError == nil && content.Active {
			continue
		}

		for _, media := range value {
			media.Active = false
		}
	}

	activeMedias := make([]cache.MediaCache, 0)
	for _, value := range medias {
		if value.Active {
			contentCache := cache.MediaCache{
				ID:                   value.ID,
				DescriptionMultiLang: nil,
				Active:               true,
			}
			activeMedias = append(activeMedias, contentCache)
		}
	}
	o.mediaCacheService.SaveAllSlice(activeMedias)
	return nil
}

func (o *mediaCacheOperations) SaveCacheMedia(ID string) (cache.MediaCache, *model.ServiceError) {
	media, err := o.mediaService.GetByID(ID)
	if err != nil || !media.Active {
		return cache.MediaCache{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Not found"}
	}

	content, err := o.contentService.GetByID(media.ContentID)
	if err != nil {
		return cache.MediaCache{}, &model.ServiceError{Code: http.StatusBadRequest, Message: "Content not found"}
	}

	mediaCache := cache.MediaCache{
		ID:                   media.ID,
		DescriptionMultiLang: nil,
	}
	mediaCache.Active = content.Active
	o.mediaCacheService.Save(mediaCache)
	return mediaCache, nil
}
