package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IGenreTypeService interface {
	Save(genreType data.GenreType) (*data.GenreType, error)
	GetAll() ([]data.GenreType, error)
	GetByID(id string) (*data.GenreType, error)
	ExistsByCode(code string) (bool, error)
	GetByCode(code string) (*data.GenreType, error)
	Update(genreType data.GenreType) (*data.GenreType, error)
	Delete(id string) error
}

type genreTypeService struct {
	genreTypeRepository repository.IGenreTypeRepository
}

func NewGenreTypeService(genreTypeRepository repository.IGenreTypeRepository) IGenreTypeService {
	return &genreTypeService{genreTypeRepository: genreTypeRepository}
}

func (s *genreTypeService) Save(genreType data.GenreType) (*data.GenreType, error) {
	return s.genreTypeRepository.Save(genreType)
}

func (s *genreTypeService) GetAll() ([]data.GenreType, error) {
	return s.genreTypeRepository.GetAll()
}

func (s *genreTypeService) GetByID(id string) (*data.GenreType, error) {
	return s.genreTypeRepository.GetByID(id)
}

func (s *genreTypeService) ExistsByCode(code string) (bool, error) {
	return s.genreTypeRepository.ExistsByCode(code)
}

func (s *genreTypeService) GetByCode(code string) (*data.GenreType, error) {
	return s.genreTypeRepository.GetByCode(code)
}

func (s *genreTypeService) Update(genreType data.GenreType) (*data.GenreType, error) {
	return s.genreTypeRepository.Update(genreType)
}

func (s *genreTypeService) Delete(id string) error {
	return s.genreTypeRepository.Delete(id)
}
