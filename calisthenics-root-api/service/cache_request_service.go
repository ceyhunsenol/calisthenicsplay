package service

import (
	"calisthenics-root-api/integration/calisthenics"
	"calisthenics-root-api/model"
)

type ICacheRequestService interface {
	GenreRefreshRequest(ID string) *model.ServiceError
	ContentRefreshRequest(ID string) *model.ServiceError
	ContentWithMediasRefreshRequest(ID string) *model.ServiceError
	MediaRefreshRequest(ID string) *model.ServiceError
	GeneralInfoRefreshRequest(ID string) *model.ServiceError
	ContentAccessRefreshRequest(ID string) *model.ServiceError
	MediaAccessRefreshRequest(ID string) *model.ServiceError
	HLSRefreshRequest(ID string) *model.ServiceError
	TranslationRefreshRequest(ID string) *model.ServiceError
}

type cacheRequestService struct {
	contentService calisthenics.ICalisthenicsContentService
}

func NewCacheRequestService(contentService calisthenics.ICalisthenicsContentService) ICacheRequestService {
	return &cacheRequestService{
		contentService: contentService,
	}
}

func (c *cacheRequestService) GenreRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "genre",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) ContentRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "content",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) ContentWithMediasRefreshRequest(ID string) *model.ServiceError {
	errorResponse := c.contentService.RefreshWithMedias(ID)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) MediaRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "media",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) GeneralInfoRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "general_info",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) ContentAccessRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "content_access",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) MediaAccessRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "media_access",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) HLSRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "hls",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}

func (c *cacheRequestService) TranslationRefreshRequest(ID string) *model.ServiceError {
	request := calisthenics.RefreshRequest{
		CacheType: "hls",
		ID:        ID,
	}
	errorResponse := c.contentService.Refresh(request)
	if errorResponse != nil {
		return &model.ServiceError{
			Message: errorResponse.Message,
		}
	}
	return nil
}
