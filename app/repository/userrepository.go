package repository

import (
	"context"

	"time"

	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/user"
	"github.com/go-redis/redis/v8"
)

type UserRepository struct {
	SQLHandler   *gen.Client
	Ctx          context.Context
	RedisHandler *redis.Client
}

func (ur *UserRepository) Insert(User *rules.RegisterValidation) (user *gen.User, err error) {
	create, err := ur.SQLHandler.User.Create().SetName(User.Name).SetUsername(User.Username).SetPassword(User.Password).SetRoleID(4).Save(ur.Ctx)

	return create, err
}

func (ur *UserRepository) CountByUsername(input string) (res int64) {
	check, _ := ur.SQLHandler.User.Query().Where(user.Username(input)).Count(ur.Ctx)

	return int64(check)
}

func (ur *UserRepository) FindByUsername(input string) (res *gen.User, err error) {
	username, err := ur.SQLHandler.User.Query().Where(user.Username(input)).Only(ur.Ctx)
	return username, err
}

func (ur *UserRepository) FindById(input uint64) (res *gen.User, err error) {
	user, err := ur.SQLHandler.User.Query().Where(user.ID(int(input))).Only(ur.Ctx)
	return user, err
}

func (ur *UserRepository) InsertRedis(key string, value interface{}, expires time.Duration) error {
	return ur.RedisHandler.Set(ur.Ctx, key, value, expires).Err()
}

func (ur *UserRepository) GettRedis(key string) (res string, err error) {
	return ur.RedisHandler.Get(ur.Ctx, key).Result()
}
