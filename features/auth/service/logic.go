package service

import (
	"errors"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/auth"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
	"github.com/go-playground/validator/v10"
)

type authService struct {
	authData auth.AuthDataInterface
	validate *validator.Validate
}

func New(repo auth.AuthDataInterface) auth.AuthServiceInterface {
	return &authService{
		authData: repo,
		validate: validator.New(),
	}
}

// Login implements auth.AuthServiceInterface.
func (service *authService) Login(email string, password string) (string, error) {
	dataLogin := auth.UserCore{
		Email:    email,
		Password: password,
	}
	errValidate := validation.ValidateStruct(service.validate, dataLogin)
	if errValidate != nil {
		return "", errValidate
	}

	data, err := service.authData.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if encrypt.CheckPasswordHash(data.Password, password) {
		token, err := service.authData.CreateToken(int(data.Id), data.Role)
		if err != nil {
			return "", errors.New(config.JWT_FailedCreateToken)
		}
		return token, nil
	}
	return "", errors.New(config.ERR_AuthWrongCredentials)
}
