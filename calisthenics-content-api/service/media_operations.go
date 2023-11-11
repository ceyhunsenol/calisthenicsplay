package service

import (
	"calisthenics-content-api/cache"
)

type IMediaOperations interface {
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
