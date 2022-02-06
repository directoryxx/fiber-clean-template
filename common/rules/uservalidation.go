package rules

type RegisterValidation struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}

type LoginValidation struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
