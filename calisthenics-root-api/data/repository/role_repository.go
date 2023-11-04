package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IRoleRepository interface {
	Save(user data.Role) (*data.Role, error)
	GetAll() ([]data.Role, error)
	GetByID(id uint) (*data.Role, error)
	GetByCode(code string) (*data.Role, error)
	ExistsByCode(code string) (bool, error)
	Update(role data.Role) (*data.Role, error)
	Delete(id uint) error
}

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) Save(role data.Role) (*data.Role, error) {
	result := r.DB.Create(&role)
	return &role, result.Error
}

func (r *roleRepository) GetAll() ([]data.Role, error) {
	var roles []data.Role
	result := r.DB.Find(&roles)
	return roles, result.Error
}

func (r *roleRepository) GetByID(id uint) (*data.Role, error) {
	var role data.Role
	result := r.DB.Where("id = ?", id).Preload("Privileges").First(&role)
	return &role, result.Error
}

func (r *roleRepository) GetByCode(code string) (*data.Role, error) {
	var role data.Role
	result := r.DB.Where("code = ?", code).First(&role)
	return &role, result.Error
}

func (r *roleRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.Role{}).Where("code = ?", code).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *roleRepository) Update(role data.Role) (*data.Role, error) {
	result := r.DB.Updates(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

func (r *roleRepository) Delete(id uint) error {
	return r.DB.Delete(&data.Role{}, id).Error
}
