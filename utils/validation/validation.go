package validation

import (
	"fmt"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Message string
	Errors  []string
}

func (err ValidationError) Error() string {
	return err.Message
}

func messageForErrTag(msg validator.FieldError) string {
	switch msg.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return msg.Error()
}

func ValidateStruct(validate *validator.Validate, data any) error {
	var errors []string
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field: %s - %s", err.Field(), messageForErrTag(err)))
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
