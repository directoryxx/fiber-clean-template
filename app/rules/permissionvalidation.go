package rules

type PermissionUpdate struct {
	Permission []UpdateRequestPermission `json:"permission" validate:"required"`
}

type UpdateRequestPermission struct {
	Page     string `json:"page" validate:"required"`
	Resource string `json:"resource" validate:"required,resource"`
}
