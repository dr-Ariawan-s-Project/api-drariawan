package service

import (
	"errors"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
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
func (ss *ScheduleService) Create(data schedule.Core, role string) error {
	if strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_AdminAccess && strings.ToLower(role) != config.VAL_SuperAdminAccess {
		return errors.New(config.VAL_InvalidValidationAccess)
	}
	err := validation.TimesValidation(data.TimeStart, data.TimeEnd)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.CreateScheduleValidate(data)
	if err != nil {
		return errors.New(err.Error())
	}
	err = validation.TimeCheckerValidate(data.TimeStart, data.TimeEnd)
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
func (ss *ScheduleService) Update(id int, data schedule.Core, role string) error {

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
func (ss *ScheduleService) Delete(id int, role string) error {

	err := ss.qry.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.ScheduleService.
func (ss *ScheduleService) GetAll(role string) ([]schedule.Core, error) {
	// log.Println(strings.ToLower(role), config.VAL_SusterAccess)
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess {
		return []schedule.Core{}, errors.New(config.VAL_InvalidValidationAccess)
	}
	res, err := ss.qry.GetAll()
	if err != nil {
		return []schedule.Core{}, errors.New(err.Error())
	}
	return res, nil
}
