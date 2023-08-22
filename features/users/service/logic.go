package service

import (
	"errors"

	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
)

type userServ struct {
	userRepo users.Repositories
}

func NewUserServ(ur users.Repositories) *userServ {
	return &userServ{
		userRepo: ur,
	}
}

func (us *userServ) FindByUsernameOrEmail(param string) (*users.Users, error) {
	if param != "" {
		return &users.Users{}, errors.New("param not found")
	}
	return us.userRepo.GetByUsername(param)
}

func (us *userServ) Insert(data *users.Users) (int, error) {
	res, err := us.userRepo.Insert(*data)
	if err != nil {
		return 0, err
	}
	return res.ID, nil
}

func (us *userServ) Update(data *users.Users, id int) error {
	_, err := us.userRepo.Update(*data, id)
	if err != nil {
		return err
	}
	return nil
}

func (us *userServ) Delete(id int) error {
	return us.userRepo.Delete(id)
}

func (us *userServ) FindById(id int) (*users.Users, error) {
	res, err := us.userRepo.GetByID(id)
	if err != nil {
		return &users.Users{}, err
	}
	return res, nil
}
func (us *userServ) FindAll(page int, rp int, param string) []*users.Users {
	res, _ := us.userRepo.Select(param, rp, page)
	return res
}
