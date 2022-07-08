package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type ValidationError struct {
	Field string
	Tag   string
	Value string
}

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
