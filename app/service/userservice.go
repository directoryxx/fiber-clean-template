package service

import (
	"github.com/directoryxx/fiber-clean-template/app/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us UserService) CreateUser(User map[string]string) error {
	_, err := us.UserRepository.Insert(User)

	return err
}
