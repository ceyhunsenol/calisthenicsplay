package service

import (
	"calisthenics-content-api/cache"
	"calisthenics-content-api/model"
)

type IGenreOperations interface {
	GetGenres(request model.GenreRequest) []model.GenreModel
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

func (o *genreOperations) GetGenres(request model.GenreRequest) []model.GenreModel {
	genreCaches := o.genreCacheService.GetAllByType(request.Type)
	models := make([]model.GenreModel, 0)
	for _, c := range genreCaches {
		if c.Section == request.Section {
			genreModel := model.GenreModel{
				ID:          c.ID,
				Type:        c.Type,
				Code:        "",
				Description: "",
				Section:     c.Section,
				Active:      true,
				Contents:    nil,
			}
			models = append(models, genreModel)
		}
	}
	return models
}
