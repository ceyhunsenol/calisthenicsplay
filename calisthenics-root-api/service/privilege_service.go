package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IPrivilegeService interface {
	Save(user data.Privilege) (*data.Privilege, error)
	GetAll() ([]data.Privilege, error)
	GetByID(id string) (*data.Privilege, error)
	ExistsByCode(code string) (bool, error)
	GetByCode(code string) (*data.Privilege, error)
	Update(privilege data.Privilege) (*data.Privilege, error)
	Delete(id string) error
}

type privilegeService struct {
	privilegeRepository repository.IPrivilegeRepository
}

func NewPrivilegeService(privilegeRepository repository.IPrivilegeRepository) IPrivilegeService {
	return &privilegeService{privilegeRepository: privilegeRepository}
}

func (s *privilegeService) Save(user data.Privilege) (*data.Privilege, error) {
	return s.privilegeRepository.Save(user)
}

func (s *privilegeService) GetAll() ([]data.Privilege, error) {
	return s.privilegeRepository.GetAll()
}

func (s *privilegeService) GetByID(id string) (*data.Privilege, error) {
	return s.privilegeRepository.GetByID(id)
}

func (s *privilegeService) ExistsByCode(code string) (bool, error) {
	return s.privilegeRepository.ExistsByCode(code)
}

func (s *privilegeService) GetByCode(code string) (*data.Privilege, error) {
	return s.privilegeRepository.GetByCode(code)
}

func (s *privilegeService) Update(privilege data.Privilege) (*data.Privilege, error) {
	return s.privilegeRepository.Update(privilege)
}

func (s *privilegeService) Delete(id string) error {
	return s.privilegeRepository.Delete(id)
}
