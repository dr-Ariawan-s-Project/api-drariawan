package validation

import (
	"fmt"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/go-playground/validator/v10"
)

type ValidationErrorItem struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationError struct {
	Message string
	Errors  []ValidationErrorItem
}

func (err ValidationError) Error() string {
	return err.Message
}

func ValidateStruct(validate *validator.Validate, data any) error {
	errors := []ValidationErrorItem{}
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationErrorItem{
				Field: err.Field(),
				Error: fmt.Sprintf("%s|%s", err.Field(), err.ActualTag()),
			})
		}
	}
	if len(errors) > 0 {
		return ValidationError{
			Message: config.VAL_InvalidValidation,
			Errors:  errors,
		}
	}
	return nil
}
