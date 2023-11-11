package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IGenreCacheOperations interface {
	SaveCacheGenres() *model.ServiceError
	SaveCacheGenre(ID string) (cache.GenreCache, *model.ServiceError)
}

type genreCacheOperations struct {
	genreService      IGenreService
	genreCacheService cache.IGenreCacheService
}

func NewGenreCacheOperations(genreService IGenreService, genreCacheService cache.IGenreCacheService) IGenreCacheOperations {
	return &genreCacheOperations{
		genreService:      genreService,
		genreCacheService: genreCacheService,
	}
}

func (o *genreCacheOperations) SaveCacheGenres() *model.ServiceError {
	genres, err := o.genreService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Unknown error"}
	}
	activeGenres := make([]cache.GenreCache, 0)
	for _, value := range genres {
		if value.Active {
			contentCache := cache.GenreCache{
				ID:                   value.ID,
				Type:                 value.Type,
				Section:              value.Section,
				CodeMultiLang:        nil,
				DescriptionMultiLang: nil,
				Active:               true,
			}
			activeGenres = append(activeGenres, contentCache)
		}
	}

	o.genreCacheService.SaveAllSlice(activeGenres)
	return nil
}

func (o *genreCacheOperations) SaveCacheGenre(ID string) (cache.GenreCache, *model.ServiceError) {
	genre, err := o.genreService.GetByID(ID)
	if err != nil || !genre.Active {
		return cache.GenreCache{}, &model.ServiceError{Code: http.StatusNotFound, Message: "Not found"}
	}
	genreCache := cache.GenreCache{
		ID:                   genre.ID,
		CodeMultiLang:        nil,
		DescriptionMultiLang: nil,
		Type:                 genre.Type,
		Active:               true,
	}
	o.genreCacheService.Save(genreCache)
	return genreCache, nil
}
