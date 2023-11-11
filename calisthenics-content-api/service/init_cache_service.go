package service

type IInitCacheService interface {
	InitCache()
}

type initCacheService struct {
	contentCacheOperations IContentCacheOperations
	genreCacheOperations   IGenreCacheOperations
	mediaCacheOperations   IMediaCacheOperations
}

func NewInitCacheService(contentCacheOperations IContentCacheOperations,
	genreCacheOperations IGenreCacheOperations,
	mediaCacheOperations IMediaCacheOperations,
) IInitCacheService {
	return &initCacheService{
		genreCacheOperations:   genreCacheOperations,
		contentCacheOperations: contentCacheOperations,
		mediaCacheOperations:   mediaCacheOperations,
	}
}

func (c *initCacheService) InitCache() {
	c.genreCacheOperations.SaveCacheGenres()
	c.contentCacheOperations.SaveCacheContents()
	c.mediaCacheOperations.SaveCacheMedias()
}
