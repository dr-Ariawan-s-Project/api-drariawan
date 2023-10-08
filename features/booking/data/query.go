package data

import (
	"errors"
	"log"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
	"gorm.io/gorm"
)

type bookingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) booking.Data {
	return &bookingQuery{
		db: db,
	}
}

// Create implements schedule.ScheduleData.
func (bq *bookingQuery) Create(data booking.Core) error {
	qry := CoreToData(data)
	qry.DeletedAt = nil
	err := bq.db.Create(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New(err.Error())
	}
	return nil
}

// Update implements schedule.ScheduleData.
func (bq *bookingQuery) Update(id int, data booking.Core) error {
	cnv := CoreToData(data)
	cnv.DeletedAt = nil
	qry := bq.db.Model(&Bookings{}).Where("id = ?", id).Updates(&cnv)
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

// Delete implements schedule.ScheduleData.
func (bq *bookingQuery) Delete(id int) error {
	data := Bookings{}
	timeNow := time.Now()
	data.DeletedAt = &timeNow
	qry := bq.db.Model(&Bookings{}).Where("id = ?", id).Updates(&data)
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

// GetAll implements schedule.ScheduleData.
func (bq *bookingQuery) GetAll() ([]booking.Core, error) {
	qry := []Bookings{}
	err := bq.db.Where("deleted_at IS NULL").Order("created_at DESC").Preload("Patient").Preload("Schedule").Preload("Schedule.User").Find(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []booking.Core{}, errors.New(err.Error())
	}
	return DataToCoreArray(qry), nil
}

// GetByUserID implements booking.Data.
func (bq *bookingQuery) GetByUserID(userID int) ([]booking.Core, error) {
	qry := []Bookings{}
	err := bq.db.Joins("JOIN schedules ON bookings.schedule_id = schedules.id").Where("schedules.user_id = ? AND bookings.deleted_at is null", userID).Order("bookings.created_at DESC").Preload("Patient").Preload("Schedule").Preload("Schedule.User").Find(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []booking.Core{}, errors.New(err.Error())
	}
	log.Println(qry)
	return DataToCoreArray(qry), nil
}
