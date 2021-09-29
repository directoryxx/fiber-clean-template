package repository

import (
	"context"

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
	create, err := ur.SQLHandler.User.Create().SetName(User.Name).SetUsername(User.Username).SetPassword(User.Password).Save(ur.Ctx)

	return create, err
}

func (ur *UserRepository) CountByUsername(input string) (res int64) {
	check, _ := ur.SQLHandler.User.Query().Where(user.Username(input)).Count(ur.Ctx)

	return int64(check)
}
