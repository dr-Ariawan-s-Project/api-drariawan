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
	query.DeletedAt = nil
	err := uq.db.Create(&query).Error
	if err != nil {
		log.Println("query error", err.Error())
		return users.UsersCore{}, errors.New(err.Error())
	}
	return DataToCore(query), nil
}

// Update implements users.UserData.
func (uq *userQuery) Update(data users.UsersCore, id int) error {
	cnv := CoreToData(data)
	cnv.DeletedAt = nil
	qry := uq.db.Model(&Users{}).Where("id = ?", id).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements users.UserData.
func (uq *userQuery) Delete(id int) error {

	qryDelete := uq.db.Delete(&Users{}, id)
	affRow := qryDelete.RowsAffected
	if affRow <= 0 {
		log.Println("No rows affected")
		return errors.New("failed to delete user content, data not found")
	}

	return nil
}

// FindAll implements users.UserData.
func (uq *userQuery) FindAll(search string, rp int, page int) ([]users.UsersCore, error) {
	panic("unimplemented")
}

// FindByID implements users.UserData.
func (uq *userQuery) FindByID(id int) (users.UsersCore, error) {
	data := Users{}
	err := uq.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return users.UsersCore{}, errors.New(err.Error())
	}
	return DataToCore(data), nil
}

// FindByUsernameOrEmail implements users.UserData.
func (uq *userQuery) FindByUsernameOrEmail(username string) (users.UsersCore, error) {
	data := Users{}
	err := uq.db.Where("username = ?", username).First(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return users.UsersCore{}, errors.New(err.Error())
	}
	return DataToCore(data), nil
}
