package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
	Role     Role
}
