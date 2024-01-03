package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	// Create a new validator for a uuid
	validate := validator.New()

	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return false
		}
		return true
	})
	return validate
}

// ValidatorErrors func for show validation errors for each  invalid fields
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error messages for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Error()] = err.Error()
	}
	return fields
}
