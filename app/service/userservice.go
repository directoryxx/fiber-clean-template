package service

import (
	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/database/gen"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us UserService) CreateUser(User map[string]string) (user *gen.User, err error) {
	data, err := us.UserRepository.Insert(User)

	return data, err
}