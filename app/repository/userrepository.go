package repository

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/database/gen"
)

type UserRepository struct {
	SQLHandler *gen.Client
	Ctx        context.Context
}

func (ur *UserRepository) Insert(User map[string]string) (user *gen.User, err error) {
	create, err := ur.SQLHandler.User.Create().SetName(User["name"]).SetUsername(User["username"]).SetPassword(User["password"]).Save(ur.Ctx)

	return create, err
}
