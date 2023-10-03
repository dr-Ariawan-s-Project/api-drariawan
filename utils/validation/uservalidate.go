package validation

import (
	"errors"
	"fmt"

	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/go-playground/validator/v10"
)

type UserValidate struct {
	Name           string `validate:"required"`
	Email          string `validate:"required,email"`
	Role           string `validate:"required,alpha"`
	Phone          string `validate:"required,numeric"`
	Specialization string `validate:"required,alpha"`
	Password       string `validate:"required,min=5,alphanum"`
}

func CoreToRegVal(data users.UsersCore) UserValidate {
	return UserValidate{
		Name:           data.Name,
		Email:          data.Email,
		Role:           data.Role,
		Specialization: data.Specialization,
		Phone:          data.Phone,
		Password:       data.Password,
	}
}
func RegistrationValidate(data users.UsersCore) error {
	validate := validator.New()
	val := CoreToRegVal(data)
	if err := validate.Struct(val); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vlderror := ""
			if e.Field() == "Password" && e.Value() != "" {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
			if e.Value() == "" {
				vlderror = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
				return errors.New(vlderror)
			} else {
				vlderror = fmt.Sprintf("%s is not %s", e.Value(), e.Tag())
				return errors.New(vlderror)
			}
		}
	}
	return nil
}

func UpdateUserCheckValidation(data users.UsersCore) error {
	validate := validator.New()
	if data.Password == "" && data.Phone == "" && data.Specialization == "" && data.Email == "" && data.Name == "" && data.Role == "" {
		return nil
	}
	if data.Name != "" {
		err := validate.Var(data.Name, "alpha")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("name %s is not %s", data.Name, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Specialization != "" {
		err := validate.Var(data.Specialization, "alpha")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("specialization %s is not %s", data.Specialization, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Role != "" {
		err := validate.Var(data.Role, "alpha")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("role %s is not %s", data.Role, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Email != "" {
		err := validate.Var(data.Email, "email")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("email %s is not %s", data.Email, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Phone != "" {
		err := validate.Var(data.Phone, "numeric")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("phone number '%s' is not %s", data.Phone, e.Tag())
			return errors.New(vlderror)
		}
	}
	if data.Password != "" {
		err := validate.Var(data.Password, "min=3,alphanum")
		if err != nil {
			e := err.(validator.ValidationErrors)[0]
			vlderror := fmt.Sprintf("password %s is not %s", data.Password, e.Tag())
			return errors.New(vlderror)
		}
	}

	return nil
}
