package repository

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/role"
)

type RoleRepository struct {
	SQLHandler *gen.Client
	Ctx        context.Context
}

func (rr *RoleRepository) Insert(Role *rules.RoleValidation) (role *gen.Role, err error) {
	create, err := rr.SQLHandler.Role.Create().SetName(Role.Name).Save(rr.Ctx)

	return create, err
}

func (rr *RoleRepository) GetAll() (role []*gen.Role, err error) {
	role, err = rr.SQLHandler.Role.Query().All(rr.Ctx)

	return role, err
}

func (rr *RoleRepository) Update(role_id int, Role *rules.RoleValidation) (role *gen.Role, err error) {
	role, err = rr.SQLHandler.Role.UpdateOneID(role_id).SetName(role.Name).Save(rr.Ctx)
	return role, err
}

func (rr *RoleRepository) Delete(role_id int) (err error) {
	err = rr.SQLHandler.Role.DeleteOneID(role_id).Exec(rr.Ctx)
	return err
}

func (rr *RoleRepository) CountByName(input string) (res int64) {
	check, _ := rr.SQLHandler.Role.Query().Where(role.Name(input)).Count(rr.Ctx)
	return int64(check)
}

func (rr *RoleRepository) FindById(role_id int) (roleData *gen.Role, err error) {
	roleData, errRoleData := rr.SQLHandler.Role.Query().Where(role.ID(role_id)).Only(rr.Ctx)

	return roleData, errRoleData
}
