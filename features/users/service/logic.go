package service

import (
	"errors"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
	"github.com/go-playground/validator/v10"
)

type userServ struct {
	userRepo users.Data
	validate *validator.Validate
}

func New(ur users.Data) users.Service {
	return &userServ{
		userRepo: ur,
		validate: validator.New(),
	}
}

// Insert implements users.UserService.
func (us *userServ) Insert(data users.UsersCore, role string) (users.UsersCore, error) {
	if strings.ToLower(role) != config.VAL_SuperAdminAccess {
		return users.UsersCore{}, errors.New(config.VAL_Unauthorized)
	}
	err := validation.RegistrationValidate(data)
	if err != nil {
		return users.UsersCore{}, errors.New(err.Error())
	}
	hash, err := encrypt.HashPassword(data.Password)
	if err != nil {
		return users.UsersCore{}, errors.New(err.Error())
	}
	data.Password = hash
	res, err := us.userRepo.Insert(data)
	if err != nil {
		return users.UsersCore{}, errors.New(err.Error())
	}
	return res, nil
}

// Update implements users.UserService.
func (us *userServ) Update(data users.UsersCore, id int) error {
	err := validation.UpdateUserCheckValidation(data)
	if err != nil {
		return errors.New(err.Error())
	}
	if data.Password != "" {
		hash, err := encrypt.HashPassword(data.Password)
		if err != nil {
			return errors.New(err.Error())
		}
		data.Password = hash
	}
	err = us.userRepo.Update(data, id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements users.UserService.
func (us *userServ) Delete(id int) error {
	err := us.userRepo.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// FindAll implements users.UserService.
func (us *userServ) FindAll(search string, rp int, page int) ([]users.UsersCore, error) {
	res, err := us.userRepo.FindAll(search, rp, page)
	if err != nil {
		return []users.UsersCore{}, errors.New(err.Error())
	}
	return res, nil
}

// FindById implements users.UserService.
func (us *userServ) FindById(id int) (users.UsersCore, error) {
	res, err := us.userRepo.FindByID(id)
	if err != nil {

		return users.UsersCore{}, errors.New(err.Error())
	}
	return res, nil
}
