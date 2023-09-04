package validation

import (
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
	errors := []ValidationErrorItem{}
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationErrorItem{
				Field: err.Field(),
				Error: messageForErrTag(err),
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
