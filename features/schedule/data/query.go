package data

import (
	"errors"
	"log"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
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

// Create implements schedule.ScheduleData.
func (sq *scheduleQuery) Create(data schedule.Core) error {
	qry := CoreToData(data)
	qry.DeletedAt = nil
	err := sq.db.Create(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
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
		log.Println("update user query error", err.Error())
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
		log.Println("update user query error", err.Error())
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.ScheduleData.
func (sq *scheduleQuery) GetAll() ([]schedule.Core, error) {
	qry := []Schedules{}
	err := sq.db.Where("deleted_at is null").Find(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []schedule.Core{}, errors.New(err.Error())
	}
	return DataToCoreArray(qry), nil
}
