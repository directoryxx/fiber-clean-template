package repository

import (
	"context"

	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/infrastructure"
	"github.com/directoryxx/fiber-clean-template/app/utils/pagination"
)

type PhoneBookRepository struct {
	// SQLHandler *gen.Client
	Ctx context.Context
}

func (pbr *PhoneBookRepository) InsertPhoneBook(Role *domain.PhoneBook) (role *domain.PhoneBook, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Create(&Role)

	// defer conn.Close()
	return Role, err
}

func (pbr *PhoneBookRepository) GetAllPhoneBook(page int, limit int) (role *[]domain.PhoneBook, err error) {
	var Role *[]domain.PhoneBook
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

func (pbr *PhoneBookRepository) UpdatePhoneBook(role_id int, Role *domain.PhoneBook) (role *domain.PhoneBook, err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Model(&domain.PhoneBook{}).Where("id = ?", role_id).Updates(Role)
	// conn.PhoneBook.UpdateOneID(role_id).SetName(Role.Name).Save(rr.Ctx)

	return Role, err
}

func (pbr *PhoneBookRepository) DeletePhoneBook(role_id int) (err error) {
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Delete(&domain.PhoneBook{}, role_id)
	return err
}

func (pbr *PhoneBookRepository) CountByNamePhoneBook(input string) (res int64) {
	var count int64
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}

	conn.Model(&domain.PhoneBook{}).Where("name = ?", input).Count(&count)
	return count
}

func (pbr *PhoneBookRepository) FindByIdPhoneBook(role_id int) (roleData *domain.PhoneBook, err error) {
	var role *domain.PhoneBook
	conn, err := infrastructure.Open()
	if err != nil {
		panic(err)
	}
	conn.Model(&domain.PhoneBook{}).Where("id = ?", role_id).First(&role)

	return role, err
}
