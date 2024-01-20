package service

import (
	"errors"
	"testing"

	"github.com/dr-ariawan-s-project/api-drariawan/features/auth"
	"github.com/dr-ariawan-s-project/api-drariawan/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	authRepo := new(mocks.AuthData)

	t.Run("1. error validate", func(t *testing.T) {

		authService := New(authRepo)
		response, token, err := authService.Login("", "qwerty")
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Nil(t, response)
	})
	t.Run("2. error GetUserByEmail", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		authRepo.On("GetUserByEmail", email).Return(nil, errors.New("error")).Once()
		authService := New(authRepo)
		response, token, err := authService.Login(email, password)
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Nil(t, response)
	})
	t.Run("3. error create token login", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		returnData := auth.UserCore{
			Id:       1,
			Email:    "test@mail.com",
			Password: "$2a$04$s3mXRcaPepFFiV8GcPcrqOT5znxxXacvkuD1SzcYNMJ6Z6MUD64g.",
			Name:     "test user",
			Role:     "dokter",
		}
		authRepo.On("GetUserByEmail", email).Return(&returnData, nil).Once()
		authRepo.On("CreateToken", int(returnData.Id), returnData.Role).Return("", errors.New("error create token")).Once()
		authService := New(authRepo)
		_, _, err := authService.Login(email, password)
		assert.Error(t, err)
	})
	t.Run("4. success login", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		returnData := auth.UserCore{
			Id:       1,
			Email:    "test@mail.com",
			Password: "$2a$04$s3mXRcaPepFFiV8GcPcrqOT5znxxXacvkuD1SzcYNMJ6Z6MUD64g.",
			Name:     "test user",
			Role:     "dokter",
		}
		authRepo.On("GetUserByEmail", email).Return(&returnData, nil).Once()
		authRepo.On("CreateToken", int(returnData.Id), returnData.Role).Return("token", nil).Once()
		authService := New(authRepo)
		response, token, err := authService.Login(email, password)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Email, response.Email)
		assert.Equal(t, "token", token)
	})
	t.Run("5. error password didnot match", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty1"
		returnData := auth.UserCore{
			Id:       1,
			Email:    "test@mail.com",
			Password: "$2a$04$s3mXRcaPepFFiV8GcPcrqOT5znxxXacvkuD1SzcYNMJ6Z6MUD64g.",
			Name:     "test user",
			Role:     "dokter",
		}
		authRepo.On("GetUserByEmail", email).Return(&returnData, nil).Once()
		// authRepo.On("CreateToken", int(returnData.Id), returnData.Role).Return("", errors.New("error create token")).Once()
		authService := New(authRepo)
		_, _, err := authService.Login(email, password)
		assert.Error(t, err)
	})
}

func TestLoginPatient(t *testing.T) {
	authRepo := new(mocks.AuthData)

	t.Run("1. error validate", func(t *testing.T) {

		authService := New(authRepo)
		response, token, err := authService.LoginPatient("", "qwerty")
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Nil(t, response)
	})

	t.Run("2. error GetPatientByEmail", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		authRepo.On("GetPatientByEmail", email).Return(nil, errors.New("error")).Once()
		authService := New(authRepo)
		response, token, err := authService.LoginPatient(email, password)
		assert.Error(t, err)
		assert.Equal(t, "", token)
		assert.Nil(t, response)
	})
	t.Run("3. error create token login", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		returnData := auth.PatientCore{
			Id:       "PATIENT0001",
			Email:    "test@mail.com",
			Password: "$2a$04$s3mXRcaPepFFiV8GcPcrqOT5znxxXacvkuD1SzcYNMJ6Z6MUD64g.",
			Name:     "test user",
			Phone:    "0812345",
		}
		authRepo.On("GetPatientByEmail", email).Return(&returnData, nil).Once()
		authRepo.On("CreateToken", returnData.Id, "patient").Return("", errors.New("error create token")).Once()
		authService := New(authRepo)
		_, _, err := authService.LoginPatient(email, password)
		assert.Error(t, err)
	})
	t.Run("4. success login", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		returnData := auth.PatientCore{
			Id:       "PATIENT0001",
			Email:    "test@mail.com",
			Password: "$2a$04$s3mXRcaPepFFiV8GcPcrqOT5znxxXacvkuD1SzcYNMJ6Z6MUD64g.",
			Name:     "test user",
			Phone:    "0812345",
		}
		authRepo.On("GetPatientByEmail", email).Return(&returnData, nil).Once()
		authRepo.On("CreateToken", returnData.Id, "patient").Return("token", nil).Once()
		authService := New(authRepo)
		response, token, err := authService.LoginPatient(email, password)
		assert.NoError(t, err)
		assert.Equal(t, returnData.Email, response.Email)
		assert.Equal(t, "token", token)
	})
	t.Run("5. error password didnot match", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty1"
		returnData := auth.PatientCore{
			Id:       "PATIENT0001",
			Email:    "test@mail.com",
			Password: "$2a$04$s3mXRcaPepFFiV8GcPcrqOT5znxxXacvkuD1SzcYNMJ6Z6MUD64g.",
			Name:     "test user",
			Phone:    "0812345",
		}
		authRepo.On("GetPatientByEmail", email).Return(&returnData, nil).Once()
		// authRepo.On("CreateToken", int(returnData.Id), returnData.Role).Return("", errors.New("error create token")).Once()
		authService := New(authRepo)
		_, _, err := authService.LoginPatient(email, password)
		assert.Error(t, err)
	})
	t.Run("6. error password NOT SET", func(t *testing.T) {
		email := "test@mail.com"
		password := "qwerty"
		returnData := auth.PatientCore{
			Id:       "PATIENT0001",
			Email:    "test@mail.com",
			Password: "",
			Name:     "test user",
			Phone:    "0812345",
		}
		authRepo.On("GetPatientByEmail", email).Return(&returnData, nil).Once()
		// authRepo.On("CreateToken", int(returnData.Id), returnData.Role).Return("", errors.New("error create token")).Once()
		authService := New(authRepo)
		_, _, err := authService.LoginPatient(email, password)
		assert.Error(t, err)
	})
}
