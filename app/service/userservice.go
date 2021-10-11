package service

import (
	"time"

	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (us UserService) CreateUser(User *domain.User) (user *domain.User) {
	data := us.UserRepository.Insert(User)

	return data
}

func (us UserService) CheckUsernameCount(username string) (count int64) {
	data := us.UserRepository.CountByUsername(username)

	return data
}

func (us UserService) CheckUsername(username string) (res *domain.User) {
	data := us.UserRepository.FindByUsername(username)

	return data
}

func (us *UserService) InsertToken(key string, value interface{}, expires time.Duration) error {
	return us.UserRepository.InsertRedis(key, value, expires)
}

func (us *UserService) FetchToken(key string) (res string, err error) {
	return us.UserRepository.GettRedis(key)
}

func (us *UserService) CurrentUser(input uint64) (res *domain.User) {
	return us.UserRepository.FindByIdWithRelation(input)
}
