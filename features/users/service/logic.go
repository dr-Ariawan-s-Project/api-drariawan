package service

import (
	"errors"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
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

// for pagination
// GetPagination implements users.Service.
func (us *userServ) GetPagination(search string, role string, page int, perPage int) (map[string]any, error) {
	totalRows, err := us.userRepo.CountByFilter(search, role)
	response := map[string]any{
		"page":          0,
		"limit":         0,
		"total_pages":   0,
		"total_records": 0,
	}
	if err != nil {
		return response, err
	}
	paginationRes := helpers.CountPagination(totalRows, page, perPage)
	response["page"] = paginationRes.Page
	response["limit"] = paginationRes.Limit
	response["total_pages"] = paginationRes.TotalPages
	response["total_records"] = paginationRes.TotalRecords
	return response, nil
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
func (us *userServ) FindAll(search string, role string, rp int, page int) ([]users.UsersCore, error) {
	if rp == 0 {
		rp = 10
	}
	if page == 0 {
		page = 1
	}

	if role != "" && role != config.VAL_AdminAccess && role != config.VAL_SuperAdminAccess && role != config.VAL_DokterAccess && role != config.VAL_SusterAccess && role != config.VAL_PatientAccess {
		return nil, errors.New("[validation] invalid role")
	}
	res, err := us.userRepo.FindAll(search, role, rp, page)
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
