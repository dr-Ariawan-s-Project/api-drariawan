package service

import (
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
)

type userServ struct {
	userRepo users.UserData
}

func NewUserServ(ur users.UserData) users.UserService {
	return &userServ{
		userRepo: ur,
	}
}

// Insert implements users.UserService.
func (us *userServ) Insert(data users.UsersCore) (int, error) {
	panic("unimplemented")
}

// Update implements users.UserService.
func (us *userServ) Update(data users.UsersCore, id int) error {
	panic("unimplemented")
}

// Delete implements users.UserService.
func (us *userServ) Delete(id int) error {
	panic("unimplemented")
}

// func (us *userServ) FindByUsernameOrEmail(param string) (*users.UsersCore, error) {
// 	if param != "" {
// 		return &users.UsersCore{}, errors.New("param not found")
// 	}
// 	return us.userRepo.GetByUsername(param)
// }

// func (us *userServ) Insert(data *users.UsersCore) (int, error) {
// 	res, err := us.userRepo.Insert(*data)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return res.ID, nil
// }

// func (us *userServ) Update(data *users.UsersCore, id int) error {
// 	_, err := us.userRepo.Update(*data, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (us *userServ) Delete(id int) error {
// 	return us.userRepo.Delete(id)
// }

// func (us *userServ) FindById(id int) (*users.UsersCore, error) {
// 	res, err := us.userRepo.GetByID(id)
// 	if err != nil {
// 		return &users.UsersCore{}, err
// 	}
// 	return res, nil
// }
// func (us *userServ) FindAll(page int, rp int, param string) []*users.UsersCore {
// 	res, _ := us.userRepo.Select(param, rp, page)
// 	return res
// }
