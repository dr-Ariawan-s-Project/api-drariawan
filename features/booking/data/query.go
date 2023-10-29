package data

import (
	"errors"
	"log"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
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
	qry.State = "confirmed"
	qry.ID = helpers.UUIDGenerate()
	qry.BookingCode = helpers.RandomStringAlphabetNumeric()

	// CEK APAKAH SUDAH PERNAH BOOKING
	cnv1 := Bookings{}
	err := bq.db.Where("booking_date = ? AND schedule_id = ? ", qry.BookingDate, qry.ScheduleId).First(&cnv1).Error
	if err == nil {
		return errors.New(config.DB_ERR_DUPLICATE_BOOKING)
	}

	// CEK APAKAH DALAM SATU MINGGU SUDAH PERNAH BOOKING
	sevenDaysLaterStr, sevenDaysAgoStr, err := validation.SevenDayLimitVal(qry.BookingDate)
	if err != nil {
		return errors.New(err.Error())
	}
	cnv2 := Bookings{}
	err = bq.db.Where("booking_date >= ? AND booking_date <= ? AND patient_id = ?", sevenDaysAgoStr, sevenDaysLaterStr, qry.PatientId).First(&cnv2).Error
	if err == nil {
		return errors.New(config.DB_ERR_LIMIT_BOOKING_SEVDAY)
	}

	qry.DeletedAt = nil
	err = bq.db.Create(&qry).Error
	if err != nil {
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
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.ScheduleData.
func (bq *bookingQuery) GetAll() ([]booking.Core, error) {
	qry := []Bookings{}
	err := bq.db.Where("deleted_at IS NULL").Order("created_at DESC").Preload("Patient").Preload("Schedule").Preload("Schedule.User").Find(&qry).Error
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return DataToCoreArray(qry), nil
}

// GetByUserID implements booking.Data.
func (bq *bookingQuery) GetByUserID(userID int) ([]booking.Core, error) {
	qry := []Bookings{}
	err := bq.db.Joins("JOIN schedules ON bookings.schedule_id = schedules.id").Where("schedules.user_id = ? AND bookings.deleted_at is null", userID).Order("bookings.created_at DESC").Preload("Patient").Preload("Schedule").Preload("Schedule.User").Find(&qry).Error
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	log.Println(qry)
	return DataToCoreArray(qry), nil
}
