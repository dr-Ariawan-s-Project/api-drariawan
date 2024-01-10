package data

import (
	"errors"
	"log"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Data {
	return &userQuery{
		db: db,
	}
}

// for pagination
// CountByFilter implements users.Data.
func (uq *userQuery) CountByFilter(search string, role string) (int64, error) {
	var countAttemp int64
	tx := uq.db.Model(&Users{})
	if search != "" {
		tx.Where("name like ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if role != "" {
		tx.Where("role = ?", role)
	}
	tx.Count(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return countAttemp, nil
}

// Insert implements users.UserData.
func (uq *userQuery) Insert(data users.UsersCore) (users.UsersCore, error) {
	query := CoreToData(data)
	log.Println(query)
	query.State = "active"
	query.DeletedAt = nil
	query.UrlPicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"
	err := uq.db.Create(&query).Error
	if err != nil {
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
	data := Users{}
	data.State = "deactive"
	timeNow := time.Now()
	data.DeletedAt = &timeNow
	qry := uq.db.Model(&Users{}).Where("id = ?", id).Updates(&data)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("no data changed")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return errors.New(err.Error())
	}
	return nil
}

// FindAll implements users.UserData.
func (uq *userQuery) FindAll(search string, role string, rp int, page int) ([]users.UsersCore, error) {
	data := []Users{}
	offset := (page - 1) * rp
	txSelect := uq.db.Where("state = ? AND deleted_at is null", "active")
	if search != "" {
		txSelect.Where("name like ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if role != "" {
		txSelect.Where("role = ?", role)
	}
	txSelect.Offset(offset).Limit(rp).Find(&data)
	if txSelect.Error != nil {
		return []users.UsersCore{}, errors.New(txSelect.Error.Error())
	}
	return DataToCoreArray(data), nil
}

// FindByID implements users.UserData.
func (uq *userQuery) FindByID(id int) (users.UsersCore, error) {
	data := Users{}
	err := uq.db.Where("id = ? AND deleted_at is null AND state = ?", id, "active").First(&data).Error
	if err != nil {
		return users.UsersCore{}, errors.New(err.Error())
	}
	return DataToCore(data), nil
}
