package service

import (
	"errors"

	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
)

type ScheduleService struct {
	qry schedule.ScheduleData
}

func New(sd schedule.ScheduleData) schedule.ScheduleService {
	return &ScheduleService{
		qry: sd,
	}
}

// Create implements schedule.ScheduleService.
func (ss *ScheduleService) Create(data schedule.Core) error {
	err := validation.TimesValidation(data.TimeStart, data.TimeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.CreateValidate(data)
	if err != nil {
		return errors.New(err.Error())
	}
	err = ss.qry.Create(data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Update implements schedule.ScheduleService.
func (ss *ScheduleService) Update(id int, data schedule.Core) error {
	err := validation.TimesValidation(data.TimeStart, data.TimeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.UpdateScheduleCheckValidation(data)
	if err != nil {
		return errors.New(err.Error())
	}
	err = ss.qry.Update(id, data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements schedule.ScheduleService.
func (ss *ScheduleService) Delete(id int) error {
	err := ss.qry.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.ScheduleService.
func (ss *ScheduleService) GetAll() ([]schedule.Core, error) {
	res, err := ss.qry.GetAll()
	if err != nil {
		return []schedule.Core{}, errors.New(err.Error())
	}
	return res, nil
}
