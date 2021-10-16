package domain

import "gorm.io/gorm"

type PhoneBook struct {
	gorm.Model
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Company  string `json:"company"`
	Position string `json:"position"`
}
