package service

import (
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
)

type IRoleService interface {
	Save(user data.Role) (*data.Role, error)
	GetAll() ([]data.Role, error)
	GetByID(id uint) (*data.Role, error)
	ExistsByCode(code string) (bool, error)
	GetByCode(code string) (*data.Role, error)
	Update(role data.Role) (*data.Role, error)
	Delete(id uint) error
}

type roleService struct {
	roleRepository repository.IRoleRepository
}

func NewRoleService(roleRepository repository.IRoleRepository) IRoleService {
	return &roleService{roleRepository: roleRepository}
}

func (s *roleService) Save(user data.Role) (*data.Role, error) {
	return s.roleRepository.Save(user)
}

func (s *roleService) GetAll() ([]data.Role, error) {
	return s.roleRepository.GetAll()
}

func (s *roleService) GetByID(id uint) (*data.Role, error) {
	return s.roleRepository.GetByID(id)
}

func (s *roleService) ExistsByCode(code string) (bool, error) {
	return s.roleRepository.ExistsByCode(code)
}

func (s *roleService) GetByCode(code string) (*data.Role, error) {
	return s.roleRepository.GetByCode(code)
}

func (s *roleService) Update(role data.Role) (*data.Role, error) {
	return s.roleRepository.Update(role)
}

func (s *roleService) Delete(id uint) error {
	return s.roleRepository.Delete(id)
}
