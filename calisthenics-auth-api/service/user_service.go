package service

import (
	"calisthenics-auth-api/data"
	"calisthenics-auth-api/data/repository"
)

type IUserService interface {
	Save(user data.User) (*data.User, error)
	Update(user data.User) (*data.User, error)
	GetById(id string) (*data.User, error)
	GetByEmail(email string) (data.User, error)
	GetByUsername(username string) (data.User, error)
	EmailExists(email string) (bool, error)
	UsernameExists(username string) (bool, error)
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetById(id string) (*data.User, error) {
	return s.userRepository.GetById(id)
}

func (s *userService) GetByEmail(email string) (data.User, error) {
	return s.userRepository.GetByEmail(email)
}

func (s *userService) GetByUsername(username string) (data.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s *userService) Save(user data.User) (*data.User, error) {
	return s.userRepository.Save(user)
}

func (s *userService) Update(user data.User) (*data.User, error) {
	return s.userRepository.Update(user)
}

func (s *userService) EmailExists(email string) (bool, error) {
	return s.userRepository.EmailExists(email)
}

func (s *userService) UsernameExists(username string) (bool, error) {
	return s.userRepository.UsernameExists(username)
}
