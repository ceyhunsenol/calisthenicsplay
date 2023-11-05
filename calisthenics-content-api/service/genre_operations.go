package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
	"net/http"
)

type IGenreOperations interface {
	SaveCacheGenres() error
	SaveCacheGenre(ID string) (cache.GenreCache, error)
}

type genreOperations struct {
	genreService      IGenreService
	genreCacheService cache.IGenreCacheService
}

func NewGenreOperations(genreService IGenreService, genreCacheService cache.IGenreCacheService) IGenreOperations {
	return &genreOperations{
		genreService:      genreService,
		genreCacheService: genreCacheService,
	}
}

func (o *genreOperations) SaveCacheGenres() error {
	genres, err := o.genreService.GetAll()
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "General error"}
	}
	activeGenres := make([]cache.GenreCache, 0)
	for _, value := range genres {
		if value.Active {
			contentCache := cache.GenreCache{
				ID:                   value.ID,
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

func (o *genreOperations) SaveCacheGenre(ID string) (cache.GenreCache, error) {
	genre, err := o.genreService.GetByID(ID)
	if err != nil {
		return cache.GenreCache{}, &model.ServiceError{Code: http.StatusInternalServerError, Message: "General error"}
	}
	genreCache := cache.GenreCache{
		ID:                   genre.ID,
		CodeMultiLang:        nil,
		DescriptionMultiLang: nil,
		Active:               true,
	}
	o.genreCacheService.Save(genreCache)
	return genreCache, nil
}
