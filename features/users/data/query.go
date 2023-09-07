package data

import (
	"errors"
	"log"

	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserData {
	return &userQuery{
		db: db,
	}
}

// Insert implements users.UserData.
func (uq *userQuery) Insert(data users.UsersCore) (users.UsersCore, error) {
	query := CoreToData(data)
	log.Println(query)
	query.State = "active"
	err := uq.db.Create(&query).Error
	if err != nil {
		log.Println("query error", err.Error())
		return users.UsersCore{}, errors.New(err.Error())
	}
	return DataToCore(query), nil
}

// Update implements users.UserData.
func (uq *userQuery) Update(data users.UsersCore, id int) (users.UsersCore, error) {
	panic("unimplemented")
}

// Delete implements users.UserData.
func (uq *userQuery) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements users.UserData.
func (uq *userQuery) FindAll(search string, rp int, page int) ([]users.UsersCore, error) {
	panic("unimplemented")
}

// FindByID implements users.UserData.
func (uq *userQuery) FindByID(id int) (users.UsersCore, error) {
	panic("unimplemented")
}

// FindByUsernameOrEmail implements users.UserData.
func (uq *userQuery) FindByUsernameOrEmail(username string) (users.UsersCore, error) {
	panic("unimplemented")
}
