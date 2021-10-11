package repository

import (
	"context"

	"time"

	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
)

type UserRepository struct {
	// SQLHandler   *gen.Client
	Ctx context.Context
}

func (ur *UserRepository) Insert(User *domain.User) (user *domain.User) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}

	conn.Create(User)
	// defer conn.Close()
	return User
}

func (ur *UserRepository) CountByUsername(input string) (res int64) {
	var count int64
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}

	conn.Model(&domain.User{}).Where("username = ?", input).Count(&count)
	// defer conn.Close()
	return count
}

func (ur *UserRepository) FindByUsername(input string) (res *domain.User) {
	var user *domain.User
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}

	conn.Model(&domain.User{}).Where("username = ?", input).First(&user)

	return user
}

func (ur *UserRepository) FindById(input uint64) (res *[]domain.User) {
	var user *[]domain.User
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Model(&user).Where("id = ?", input).Find(&user)
	// defer conn.Close()
	return user
}

func (ur *UserRepository) FindByIdWithRelation(input uint64) (res *domain.User) {
	var user *domain.User
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Model(&user).Where("id = ?", input).Preload("Role").First(&user)

	return user
}

func (ur *UserRepository) InsertRedis(key string, value interface{}, expires time.Duration) error {
	redisClient := infrastructure.RedisInit()
	set := redisClient.Set(ur.Ctx, key, value, expires).Err()
	defer redisClient.Close()
	return set
}

func (ur *UserRepository) GettRedis(key string) (res string, err error) {
	redisClient := infrastructure.RedisInit()
	get, err := redisClient.Get(ur.Ctx, key).Result()
	defer redisClient.Close()
	return get, err
}
