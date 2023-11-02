package data

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
	"gorm.io/gorm"
)

type scheduleQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) schedule.ScheduleData {
	return &scheduleQuery{
		db: db,
	}
}

func (sq *scheduleQuery) CheckDuplUserID(userId int) error {
	check := Schedules{}
	err := sq.db.Where("user_id = ?", userId).First(&check).Error
	if err == nil {
		return err
	} else {
		return nil
	}
}

// Create implements schedule.ScheduleData.
func (sq *scheduleQuery) Create(data schedule.Core) error {
	qry := CoreToData(data)
	qry.DeletedAt = nil
	ckData := []Schedules{}
	err := sq.db.Where("day = ? AND user_id = ?", qry.Day, qry.UserId).Find(&ckData).Error
	if err != nil {
		return errors.New(err.Error())
	}
	intInputTimeStart, _ := strconv.Atoi(strings.Replace(qry.TimeStart, ":", "", -1))
	intInputTimeEnd, _ := strconv.Atoi(strings.Replace(qry.TimeEnd, ":", "", -1))
	if len(ckData) > 0 {
		for _, val := range ckData {
			intTimeStart, _ := strconv.Atoi(strings.Replace(val.TimeStart, ":", "", -1))
			intTimeEnd, _ := strconv.Atoi(strings.Replace(val.TimeEnd, ":", "", -1))
			if intInputTimeStart >= intTimeStart && intInputTimeStart <= intTimeEnd {
				return errors.New(config.DB_ERR_DUPLICATE_SCHEDULE)
			} else if intInputTimeEnd >= intTimeStart && intInputTimeEnd <= intTimeEnd {
				return errors.New(config.DB_ERR_DUPLICATE_SCHEDULE)
			} else if intTimeStart >= intInputTimeStart && intTimeStart <= intInputTimeEnd {
				return errors.New(config.DB_ERR_DUPLICATE_SCHEDULE)
			} else if intTimeEnd >= intInputTimeStart && intTimeEnd <= intInputTimeEnd {
				return errors.New(config.DB_ERR_DUPLICATE_SCHEDULE)
			}
		}
	}

	err = sq.db.Create(&qry).Error
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Update implements schedule.ScheduleData.
func (sq *scheduleQuery) Update(id int, data schedule.Core) error {
	cnv := CoreToData(data)
	cnv.DeletedAt = nil
	qry := sq.db.Model(&Schedules{}).Where("id = ?", id).Updates(&cnv)
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
func (sq *scheduleQuery) Delete(id int) error {
	data := Schedules{}
	timeNow := time.Now()
	data.DeletedAt = &timeNow
	qry := sq.db.Model(&Schedules{}).Where("id = ?", id).Updates(&data)
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
func (sq *scheduleQuery) GetAll() ([]schedule.Core, error) {

	currentDate := time.Now().Format("2006-01-02")
	sevenDaysLaterStr, sevenDaysAgoStr, err := validation.SevenDayLimitVal(currentDate)
	if err != nil {
		return []schedule.Core{}, errors.New(err.Error())
	}
	log.Println(currentDate)
	qry := []Schedules{}
	err = sq.db.Preload("User").Preload("Booking", "booking_date >= ? AND booking_date <= ?", sevenDaysAgoStr, sevenDaysLaterStr).Where("deleted_at is null").Find(&qry).Error
	if err != nil {
		return []schedule.Core{}, errors.New(err.Error())
	}

	return DataToCoreArray(qry), nil
}
