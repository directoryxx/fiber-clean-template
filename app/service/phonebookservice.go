package service

import (
	"strconv"

	"github.com/directoryxx/fiber-clean-template/app/domain"
	"github.com/directoryxx/fiber-clean-template/app/repository"
)

type PhoneBookService struct {
	PhoneBookRepository repository.PhoneBookRepository
}

func (pbs PhoneBookService) GetAll(page string, limit string) (PhoneBookS *[]domain.PhoneBook, err int) {
	pageInt, errPage := strconv.Atoi(page)

	limitInt, errLimit := strconv.Atoi(limit)

	if page != "" && limit != "" {
		if errPage != nil || errLimit != nil {
			return nil, 1
		}
	}

	roleData, _ := pbs.PhoneBookRepository.GetAllPhoneBook(pageInt, limitInt)

	return roleData, 0
}

func (pbs PhoneBookService) CreateRole(PhoneBook *domain.PhoneBook) (user *domain.PhoneBook, err error) {
	data, err := pbs.PhoneBookRepository.InsertPhoneBook(PhoneBook)

	return data, err
}

func (pbs PhoneBookService) UpdateRole(phonebook_id int, PhoneBook *domain.PhoneBook) (user *domain.PhoneBook, err error) {
	data, err := pbs.PhoneBookRepository.UpdatePhoneBook(phonebook_id, PhoneBook)

	return data, err
}

func (pbs PhoneBookService) CheckDuplicateNameRole(name string) int64 {
	data := pbs.PhoneBookRepository.CountByNamePhoneBook(name)

	return data
}

func (pbs PhoneBookService) GetById(phonebook_id int) (user *domain.PhoneBook) {
	roleData, _ := pbs.PhoneBookRepository.FindByIdPhoneBook(phonebook_id)

	return roleData
}

func (pbs PhoneBookService) DeleteRole(phonebook_id int) error {
	deleteRole := pbs.PhoneBookRepository.DeletePhoneBook(phonebook_id)

	return deleteRole
}
