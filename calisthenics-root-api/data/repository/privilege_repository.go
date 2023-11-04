package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IPrivilegeRepository interface {
	Save(user data.Privilege) (*data.Privilege, error)
	GetAll() ([]data.Privilege, error)
	GetByID(id string) (*data.Privilege, error)
	ExistsByCode(code string) (bool, error)
	GetByCode(code string) (*data.Privilege, error)
	Update(privilege data.Privilege) (*data.Privilege, error)
	Delete(id string) error
}

type privilegeRepository struct {
	DB *gorm.DB
}

func NewPrivilegeRepository(db *gorm.DB) IPrivilegeRepository {
	return &privilegeRepository{DB: db}
}

func (r *privilegeRepository) Save(privilege data.Privilege) (*data.Privilege, error) {
	result := r.DB.Create(&privilege)
	return &privilege, result.Error
}

func (r *privilegeRepository) GetAll() ([]data.Privilege, error) {
	var privileges []data.Privilege
	result := r.DB.Find(&privileges)
	return privileges, result.Error
}

func (r *privilegeRepository) GetByID(id string) (*data.Privilege, error) {
	var privilege data.Privilege
	result := r.DB.Where("id = ?", id).First(&privilege)
	return &privilege, result.Error
}

func (r *privilegeRepository) GetByCode(code string) (*data.Privilege, error) {
	var privilege data.Privilege
	result := r.DB.Where("code = ?", code).First(&privilege)
	return &privilege, result.Error
}

func (r *privilegeRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	if err := r.DB.Model(&data.Privilege{}).Where("code = ?", code).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *privilegeRepository) Update(privilege data.Privilege) (*data.Privilege, error) {
	result := r.DB.Updates(&privilege)
	if result.Error != nil {
		return nil, result.Error
	}
	return &privilege, nil
}

func (r *privilegeRepository) Delete(id string) error {
	return r.DB.Delete(&data.Privilege{}, "id = ?", id).Error
}
