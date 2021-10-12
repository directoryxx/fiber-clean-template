package repository

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/utils/pagination"
)

type RoleRepository struct {
	// SQLHandler *gen.Client
	Ctx context.Context
}

func (rr *RoleRepository) Insert(Role *domain.Role) (role *domain.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Create(&Role)

	// defer conn.Close()
	return Role, err
}

func (rr *RoleRepository) GetAll(page int, limit int) (role *[]domain.Role, err error) {
	var Role *[]domain.Role
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	if page == 0 && limit == 0 {
		conn.Model(&Role).Preload("User").Find(&Role)
	} else {
		conn.Model(&Role).Scopes(pagination.Paginate(page, limit)).Preload("User").Find(&Role)
	}
	// defer conn.Close()
	return Role, err
}

func (rr *RoleRepository) Update(role_id int, Role *domain.Role) (role *domain.Role, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Model(&domain.Role{}).Where("id = ?", role_id).Updates(Role)
	// conn.Role.UpdateOneID(role_id).SetName(Role.Name).Save(rr.Ctx)

	return Role, err
}

func (rr *RoleRepository) Delete(role_id int) (err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Delete(&domain.Role{}, role_id)
	return err
}

func (rr *RoleRepository) CountByName(input string) (res int64) {
	var count int64
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}

	conn.Model(&domain.Role{}).Where("name = ?", input).Count(&count)
	return count
}

func (rr *RoleRepository) FindById(role_id int) (roleData *domain.Role, err error) {
	var role *domain.Role
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Model(&domain.Role{}).Where("id = ?", role_id).First(&role)

	return role, err
}
