package repository

import (
	"context"

	"time"

	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/user"
	"github.com/go-redis/redis/v8"
)

type UserRepository struct {
	// SQLHandler   *gen.Client
	Ctx          context.Context
	RedisHandler *redis.Client
}

func (ur *UserRepository) Insert(User *rules.RegisterValidation) (user *gen.User, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	create, err := conn.User.Create().SetName(User.Name).SetUsername(User.Username).SetPassword(User.Password).SetRoleID(4).Save(ur.Ctx)
	defer conn.Close()
	return create, err
}

func (ur *UserRepository) CountByUsername(input string) (res int64) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	check, _ := conn.User.Query().Where(user.Username(input)).Count(ur.Ctx)
	defer conn.Close()
	return int64(check)
}

func (ur *UserRepository) FindByUsername(input string) (res *gen.User, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	username, err := conn.User.Query().Where(user.Username(input)).Only(ur.Ctx)
	defer conn.Close()
	return username, err
}

func (ur *UserRepository) FindById(input uint64) (res *gen.User, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	user, err := conn.User.Query().Where(user.ID(int(input))).Only(ur.Ctx)
	defer conn.Close()
	return user, err
}

func (ur *UserRepository) FindByIdWithRelation(input uint64) (res *gen.User, role *gen.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	userData, err := conn.User.Query().Where(user.ID(int(input))).Only(ur.Ctx)
	role, err = conn.User.Query().Where(user.ID(int(input))).QueryRole().Only(ur.Ctx)
	defer conn.Close()
	return userData, role, err
}

func (ur *UserRepository) InsertRedis(key string, value interface{}, expires time.Duration) error {
	return ur.RedisHandler.Set(ur.Ctx, key, value, expires).Err()
}

func (ur *UserRepository) GettRedis(key string) (res string, err error) {
	return ur.RedisHandler.Get(ur.Ctx, key).Result()
}
