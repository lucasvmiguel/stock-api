package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

// Validation error might be returned but the function Validate
type ValidationError struct {
	Field string
	Tag   string
	Value string
}

// Validates if body (eg: a json body) has valid params
// Reference: github.com/go-playground/validator
func Validate(body interface{}) []ValidationError {
	err := validate.Struct(body)
	if err == nil {
		return nil
	}

	errors := []ValidationError{}
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, ValidationError{
			Field: err.Field(),
			Tag:   err.Tag(),
			Value: err.Param(),
		})
	}

	return errors
}
