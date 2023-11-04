package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/model"
	"net/http"
)

type IGenreContentOperations interface {
	AddContent(request model.GenreContentRequest) *model.ServiceError
	RemoveContent(request model.GenreContentRequest) *model.ServiceError
}

type genreContentOperations struct {
	genreService        IGenreService
	genreContentService IGenreContentService
	contentService      IContentService
}

func NewGenreContentOperations(genreService IGenreService, genreContentService IGenreContentService, contentService IContentService) IGenreContentOperations {
	return &genreContentOperations{
		genreContentService: genreContentService,
		contentService:      contentService,
		genreService:        genreService,
	}
}

func (g *genreContentOperations) AddContent(request model.GenreContentRequest) *model.ServiceError {
	content, err := g.contentService.GetByID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found."}
	}

	genre, err := g.genreService.GetByID(request.GenreID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Genre not found."}
	}

	_, err = g.genreContentService.Save(data.GenreContent{
		ContentID: content.ID,
		GenreID:   genre.ID,
	})

	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Content could not be added to the genre."}
	}

	return nil
}

func (g *genreContentOperations) RemoveContent(request model.GenreContentRequest) *model.ServiceError {
	content, err := g.contentService.GetByID(request.ContentID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Content not found."}
	}

	genre, err := g.genreService.GetByID(request.GenreID)
	if err != nil {
		return &model.ServiceError{Code: http.StatusNotFound, Message: "Genre not found."}
	}

	err = g.genreContentService.Delete(data.GenreContent{
		ContentID: content.ID,
		GenreID:   genre.ID,
	})
	if err != nil {
		return &model.ServiceError{Code: http.StatusInternalServerError, Message: "Content could not be removed from the genre."}
	}

	return nil
}
