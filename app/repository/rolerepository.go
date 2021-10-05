package repository

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/rules"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/directoryxx/fiber-clean-template/database/gen/role"
)

type RoleRepository struct {
	// SQLHandler *gen.Client
	Ctx context.Context
}

func (rr *RoleRepository) Insert(Role *rules.RoleValidation) (role *gen.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	create, err := conn.Role.Create().SetName(Role.Name).Save(rr.Ctx)
	defer conn.Close()
	return create, err
}

func (rr *RoleRepository) GetAll() (role []*gen.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	role, err = conn.Role.Query().All(rr.Ctx)
	defer conn.Close()
	return role, err
}

func (rr *RoleRepository) Update(role_id int, Role *rules.RoleValidation) (role *gen.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	roleUpdate, errUpdate := conn.Role.UpdateOneID(role_id).SetName(Role.Name).Save(rr.Ctx)
	defer conn.Close()
	return roleUpdate, errUpdate
}

func (rr *RoleRepository) Delete(role_id int) (err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	err = conn.Role.DeleteOneID(role_id).Exec(rr.Ctx)
	defer conn.Close()
	return err
}

func (rr *RoleRepository) CountByName(input string) (res int64) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	check, _ := conn.Role.Query().Where(role.Name(input)).Count(rr.Ctx)
	defer conn.Close()
	return int64(check)
}

func (rr *RoleRepository) FindById(role_id int) (roleData *gen.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	roleData, errRoleData := conn.Role.Query().Where(role.ID(role_id)).Only(rr.Ctx)
	defer conn.Close()
	return roleData, errRoleData
}
