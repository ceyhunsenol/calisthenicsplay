package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
	"gorm.io/gorm"
)

type IGenreService interface {
	Save(tx *gorm.DB, genre data.Genre) (*data.Genre, error)
	GetAll() ([]data.Genre, error)
	GetByID(id string) (*data.Genre, error)
	ExistsByCode(code string) (bool, error)
	GetByCode(code string) (*data.Genre, error)
	Update(tx *gorm.DB, genre data.Genre) (*data.Genre, error)
	Delete(tx *gorm.DB, id string) error
}

type genreService struct {
	genreRepository repository.IGenreRepository
}

func NewGenreService(genreRepository repository.IGenreRepository) IGenreService {
	return &genreService{genreRepository: genreRepository}
}

func (s *genreService) Save(tx *gorm.DB, genre data.Genre) (*data.Genre, error) {
	return s.genreRepository.Save(tx, genre)
}

func (s *genreService) GetAll() ([]data.Genre, error) {
	return s.genreRepository.GetAll()
}

func (s *genreService) GetByID(id string) (*data.Genre, error) {
	return s.genreRepository.GetByID(id)
}

func (s *genreService) ExistsByCode(code string) (bool, error) {
	return s.genreRepository.ExistsByCode(code)
}

func (s *genreService) GetByCode(code string) (*data.Genre, error) {
	return s.genreRepository.GetByCode(code)
}

func (s *genreService) Update(tx *gorm.DB, genre data.Genre) (*data.Genre, error) {
	return s.genreRepository.Update(tx, genre)
}

func (s *genreService) Delete(tx *gorm.DB, id string) error {
	return s.genreRepository.Delete(tx, id)
}
