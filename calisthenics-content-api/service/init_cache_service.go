package service

type IInitCacheService interface {
	InitCache()
}

type initCacheService struct {
	contentOperations IContentOperations
	genreOperations   IGenreOperations
	mediaOperations   IMediaOperations
}

func NewInitCacheService(contentOperations IContentOperations,
	genreOperations IGenreOperations,
	mediaOperations IMediaOperations,
) IInitCacheService {
	return &initCacheService{
		contentOperations: contentOperations,
		genreOperations:   genreOperations,
		mediaOperations:   mediaOperations,
	}
}

func (c *initCacheService) InitCache() {
	c.genreOperations.SaveCacheGenres()
	c.contentOperations.SaveCacheContents()
	c.mediaOperations.SaveCacheMedias()
}
