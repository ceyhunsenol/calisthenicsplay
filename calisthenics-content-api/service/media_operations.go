package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/data"
	"calisthenics-content-api/model"
	"calisthenics-content-api/pkg"
	"net/http"
)

type IMediaOperations interface {
	SaveCacheMedias() error
	SaveCacheMedia(ID string) (cache.MediaCache, error)
}

type mediaOperations struct {
	mediaService      IMediaService
	mediaCacheService cache.IMediaCacheService
	contentService    IContentService
}

func NewMediaOperations(mediaService IMediaService, mediaCacheService cache.IMediaCacheService, contentService IContentService) IMediaOperations {
	return &mediaOperations{
		mediaService:      mediaService,
		mediaCacheService: mediaCacheService,
		contentService:    contentService,
	}
}

func (o *mediaOperations) SaveCacheMedias() error {
	medias, err := o.mediaService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "General error"}
	}

	grouped := pkg.GroupByField(medias, func(media *data.Media) string {
		return media.ContentID
	})

	for key, value := range grouped {
		content, err := o.contentService.GetByID(key)
		if err != nil || content.Active {
			continue
		}

		if !content.Active {
			for _, media := range value {
				media.Active = false
			}
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

func (o *mediaOperations) SaveCacheMedia(ID string) (cache.MediaCache, error) {
	media, err := o.mediaService.GetByID(ID)
	if err != nil {
		return cache.MediaCache{}, &model.ServiceError{Code: http.StatusInternalServerError, Message: "General error"}
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
