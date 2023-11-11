package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IContentCacheOperations interface {
	SaveCacheContentWithMedias(ID string) *model.ServiceError
	SaveCacheContents() *model.ServiceError
	SaveCacheContent(ID string) (cache.ContentCache, *model.ServiceError)
}

type contentCacheOperations struct {
	contentService      IContentService
	contentCacheService cache.IContentCacheService
	mediaService        IMediaService
	mediaCacheService   cache.IMediaCacheService
}

func NewContentCacheOperations(contentService IContentService,
	contentCacheService cache.IContentCacheService,
	mediaService IMediaService,
	mediaCacheService cache.IMediaCacheService,
) IContentCacheOperations {
	return &contentCacheOperations{
		contentService:      contentService,
		contentCacheService: contentCacheService,
		mediaService:        mediaService,
		mediaCacheService:   mediaCacheService,
	}
}

func (o *contentCacheOperations) SaveCacheContentWithMedias(ID string) *model.ServiceError {
	content, serviceError := o.SaveCacheContent(ID)
	if serviceError != nil {
		return serviceError
	}

	medias, err := o.mediaService.GetAllByContentID(content.ID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	mediaCaches := make([]cache.MediaCache, 0)
	for _, media := range medias {
		if !media.Active {
			continue
		}
		mediaCache := cache.MediaCache{
			ID:                   media.ID,
			DescriptionMultiLang: nil,
			URL:                  "",
			Type:                 "",
			Active:               false,
			ContentID:            "",
		}
		mediaCaches = append(mediaCaches, mediaCache)
	}
	o.mediaCacheService.SaveAllSlice(mediaCaches)
	return nil
}

func (o *contentCacheOperations) SaveCacheContents() *model.ServiceError {
	contents, err := o.contentService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
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

func (o *contentCacheOperations) SaveCacheContent(ID string) (cache.ContentCache, *model.ServiceError) {
	content, err := o.contentService.GetByID(ID)
	if err != nil || !content.Active {
		return cache.ContentCache{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Not found"}
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
