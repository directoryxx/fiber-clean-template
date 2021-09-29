package service

import (
	"time"

	"github.com/directoryxx/fiber-clean-template/app/repository"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us UserService) CreateUser(User *rules.RegisterValidation) (user *gen.User, err error) {
	data, err := us.UserRepository.Insert(User)

	return data, err
}

func (us UserService) CheckUsername(username string) (count int64) {
	data := us.UserRepository.CountByUsername(username)

	return data
}

func (us *UserService) InsertToken(key string, value interface{}, expires time.Duration) error {
	return us.UserRepository.InsertRedis(key, value, expires)
}

func (us *UserService) FetchToken(key string) (res string, err error) {
	return us.UserRepository.GettRedis(key)
}
