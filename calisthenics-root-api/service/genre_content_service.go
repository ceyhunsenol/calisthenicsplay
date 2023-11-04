package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IGenreContentService interface {
	Save(genreContent data.GenreContent) (*data.GenreContent, error)
	Delete(genreContent data.GenreContent) error
}

type genreContentService struct {
	genreContentRepository repository.IGenreContentRepository
}

func NewGenreContentService(genreContentRepository repository.IGenreContentRepository) IGenreContentService {
	return &genreContentService{
		genreContentRepository: genreContentRepository,
	}
}

func (g *genreContentService) Save(genreContent data.GenreContent) (*data.GenreContent, error) {
	return g.genreContentRepository.Save(genreContent)
}

func (g *genreContentService) Delete(genreContent data.GenreContent) error {
	return g.genreContentRepository.Delete(genreContent)
}
