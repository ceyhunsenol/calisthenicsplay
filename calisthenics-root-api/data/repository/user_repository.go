package repository

import (
	"calisthenics-root-api/data"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(user data.User) (*data.User, error)
	GetById(id string) (*data.User, error)
	GetByEmail(email string) (data.User, error)
	GetByUsername(username string) (data.User, error)
	Update(user data.User) (*data.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetById(id string) (*data.User, error) {
	var user data.User
	result := r.DB.Where("id = ?", id).Preload("Roles").Preload("Roles.Privileges").First(&user)
	return &user, result.Error
}

func (r *userRepository) GetByEmail(email string) (data.User, error) {
	var user data.User
	result := r.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *userRepository) GetByUsername(username string) (data.User, error) {
	var user data.User
	result := r.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}

func (r *userRepository) Save(user data.User) (*data.User, error) {
	result := r.DB.Create(&user)
	return &user, result.Error
}

func (r *userRepository) Update(user data.User) (*data.User, error) {
	result := r.DB.Updates(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
