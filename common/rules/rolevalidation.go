package rules

type RoleValidation struct {
	Name string `json:"name" validate:"required,name"`
}
