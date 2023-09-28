package service

import (
	"errors"
	"log"

	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
)

type userServ struct {
	userRepo users.UserData
}

func New(ur users.UserData) users.UserService {
func New(ur users.UserData) users.UserService {
	return &userServ{
		userRepo: ur,
	}
}

// Insert implements users.UserService.
func (us *userServ) Insert(data users.UsersCore) (users.UsersCore, error) {
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
	if data.Password != "" {
		hash, err := encrypt.HashPassword(data.Password)
		if err != nil {
			return errors.New(err.Error())
		}
		data.Password = hash
	}
	err := us.userRepo.Update(data, id)
	if err != nil {
		return errors.New(err.Error())
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements users.UserService.
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
