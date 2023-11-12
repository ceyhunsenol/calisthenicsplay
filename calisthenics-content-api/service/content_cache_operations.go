package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IContentCacheOperations interface {
	SaveCacheContentWithMedias(ID string) *model.ServiceError
	SaveCacheContents() *model.ServiceError
	SaveCacheContent(ID string) *cache.ContentCache
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
	content := o.SaveCacheContent(ID)
	if content == nil {
		return &model.ServiceError{Code: http.StatusBadRequest, Message: "Content not found"}
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
	o.contentCacheService.RemoveAll()
	contents, err := o.contentService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}

	activeContents := make([]cache.ContentCache, 0)
	for _, value := range contents {
		if value.Active {
			//
			codeMultiLang := cache.NewMultiLangCache(value.Code)
			codeMultiLang.SetByLang("en", value.Code)
			codeMultiLang.SetByLang("base", value.Code)

			descriptionMultiLang := cache.NewMultiLangCache(value.Description)
			descriptionMultiLang.SetByLang("en", value.Description)
			descriptionMultiLang.SetByLang("base", value.Description)
			//
			contentCache := cache.ContentCache{
				ID:                    value.ID,
				CodeMultiLang:         codeMultiLang,
				DescriptionMultiLang:  descriptionMultiLang,
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

func (o *contentCacheOperations) SaveCacheContent(ID string) *cache.ContentCache {
	o.contentCacheService.Remove(ID)
	content, err := o.contentService.GetByID(ID)
	if err != nil || !content.Active {
		return nil
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
	return &contentCache
}
