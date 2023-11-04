package service

import (
	"calisthenics-content-api/data"
	"calisthenics-content-api/data/repository"
)

type IGenreService interface {
	GetAll() ([]data.Genre, error)
	GetByID(id string) (*data.Genre, error)
	GetByCode(code string) (*data.Genre, error)
}

type genreService struct {
	genreRepository repository.IGenreRepository
}

func NewGenreService(genreRepository repository.IGenreRepository) IGenreService {
	return &genreService{genreRepository: genreRepository}
}

func (s *genreService) GetAll() ([]data.Genre, error) {
	return s.genreRepository.GetAll()
}

func (s *genreService) GetByID(id string) (*data.Genre, error) {
	return s.genreRepository.GetByID(id)
}

func (s *genreService) GetByCode(code string) (*data.Genre, error) {
	return s.genreRepository.GetByCode(code)
}
