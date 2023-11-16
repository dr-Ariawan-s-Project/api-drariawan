package service

import (
	"errors"
	"strings"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
)

type bookingService struct {
	qry booking.Data
}

func New(sd booking.Data) booking.Service {
	return &bookingService{
		qry: sd,
	}
}

// Create implements booking.Service.
func (bs *bookingService) Create(data booking.Core, role string) error {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_PatientAccess {
		return errors.New(config.VAL_Unauthorized)
	}
	err := bs.qry.Create(data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Update implements schedule.bookingService.
func (bs *bookingService) Update(id int, data booking.Core, role string) error {
	err := bs.qry.Update(id, data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements schedule.bookingService.
func (bs *bookingService) Delete(id int, role string) error {
	err := bs.qry.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.bookingService.
func (bs *bookingService) GetAll(role string) ([]booking.Core, error) {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess {
		return []booking.Core{}, errors.New(config.VAL_Unauthorized)
	}
	res, err := bs.qry.GetAll()
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return res, nil
}

// GetByUserID implements booking.Service.
func (bs *bookingService) GetByUserID(userID int, role string) ([]booking.Core, error) {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess {
		return []booking.Core{}, errors.New(config.VAL_Unauthorized)
	}
	res, err := bs.qry.GetByUserID(userID)
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return res, nil
}

// GetByPatientID implements booking.Service.
func (bs *bookingService) GetByPatientID(patientID string) ([]booking.Core, error) {
	res, err := bs.qry.GetByPatientID(patientID)
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return res, nil
}
