package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IGeneralInfoService interface {
	GetAll() ([]data.GeneralInfo, error)
	GetByID(id string) (*data.GeneralInfo, error)
	GetByKey(key string) (*data.GeneralInfo, error)
}

type generalInfoService struct {
	generalInfoRepository repository.IGeneralInfoRepository
}

func NewGeneralInfoService(generalInfoRepository repository.IGeneralInfoRepository) IGeneralInfoService {
	return &generalInfoService{
		generalInfoRepository: generalInfoRepository,
	}
}

func (s *generalInfoService) GetAll() ([]data.GeneralInfo, error) {
	return s.generalInfoRepository.GetAll()
}

func (s *generalInfoService) GetByID(id string) (*data.GeneralInfo, error) {
	return s.generalInfoRepository.GetByID(id)
}

func (s *generalInfoService) GetByKey(key string) (*data.GeneralInfo, error) {
	return s.generalInfoRepository.GetByKey(key)
}
