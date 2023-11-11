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
