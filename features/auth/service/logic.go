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
func (service *authService) Login(email string, password string) (*auth.UserCore, string, error) {
	dataLogin := auth.UserCore{
		Email:    email,
		Password: password,
	}
	errValidate := validation.ValidateStruct(service.validate, dataLogin)
	if errValidate != nil {
		return nil, "", errValidate
	}

	data, err := service.authData.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if encrypt.CheckPasswordHash(data.Password, password) {
		token, err := service.authData.CreateToken(int(data.Id), data.Role)
		if err != nil {
			return nil, "", errors.New(config.JWT_FailedCreateToken)
		}
		return data, token, nil
	}
	return nil, "", errors.New(config.ERR_AuthWrongCredentials)
}

// LoginPatient implements auth.AuthServiceInterface.
func (service *authService) LoginPatient(email string, password string) (*auth.PatientCore, string, error) {
	dataLogin := auth.PatientCore{
		Email:    email,
		Password: password,
	}
	errValidate := validation.ValidateStruct(service.validate, dataLogin)
	if errValidate != nil {
		return nil, "", errValidate
	}

	data, err := service.authData.GetPatientByEmail(email)
	if err != nil {
		return nil, "", err
	}
	if data.Password == "" {
		return nil, "", errors.New(config.VAL_PasswordNotSet)
	}

	if encrypt.CheckPasswordHash(data.Password, password) {
		token, err := service.authData.CreateToken(data.Id, config.VAL_PatientAccess)
		if err != nil {
			return nil, "", errors.New(config.JWT_FailedCreateToken)
		}
		return data, token, nil
	}
	return nil, "", errors.New(config.ERR_AuthWrongCredentials)
}
