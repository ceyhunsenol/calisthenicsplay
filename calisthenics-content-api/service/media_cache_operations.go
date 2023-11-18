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
	SaveCacheMedia(ID string) interface{}
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
	o.mediaCacheService.RemoveAll()
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
			//
			codeMultiLang := cache.NewMultiLangCache(value.DescriptionCode)
			codeMultiLang.SetByLang("en", value.DescriptionCode)
			codeMultiLang.SetByLang("base", value.DescriptionCode)
			//
			mediaCache := cache.MediaCache{
				ID:                   value.ID,
				DescriptionMultiLang: codeMultiLang,
				EncodingID:           value.EncodingID,
				Active:               true,
			}
			activeMedias = append(activeMedias, mediaCache)
		}
	}
	o.mediaCacheService.SaveAllSlice(activeMedias)
	return nil
}

func (o *mediaCacheOperations) SaveCacheMedia(ID string) interface{} {
	o.mediaCacheService.Remove(ID)
	media, err := o.mediaService.GetByID(ID)
	if err != nil || !media.Active {
		return nil
	}

	content, err := o.contentService.GetByID(media.ContentID)
	if err != nil {
		return nil
	}

	mediaCache := cache.MediaCache{
		ID:                   media.ID,
		EncodingID:           media.EncodingID,
		DescriptionMultiLang: nil,
	}
	mediaCache.Active = content.Active
	o.mediaCacheService.Save(mediaCache)
	return &mediaCache
}
