package validation

import (
	"github.com/go-playground/validator/v10"
)

// ErrorValidationResponse - Standarize the response for validation
type ErrorValidationResponse struct {
	FailedField string
	Tag         string
	Value       string
	Message     string
}

func ValidateRequest(err error, message map[string]string) []*ErrorValidationResponse {
	var errors []*ErrorValidationResponse
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidationResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Message = message[err.Tag()]
			errors = append(errors, &element)
		}
	}
	return errors
}

// ValidateStruct - Validate Input for all usecase
func ValidateStruct(class interface{}) []*ErrorValidationResponse {
	var errors []*ErrorValidationResponse
	validate := validator.New()

	err := validate.Struct(class)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidationResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}