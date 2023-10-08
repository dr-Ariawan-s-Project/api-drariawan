package service

import (
	"errors"

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
func (bs *bookingService) Create(data booking.Core) error {
	err := bs.qry.Create(data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Update implements schedule.bookingService.
func (bs *bookingService) Update(id int, data booking.Core) error {
	err := bs.qry.Update(id, data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements schedule.bookingService.
func (bs *bookingService) Delete(id int) error {
	err := bs.qry.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetAll implements schedule.bookingService.
func (bs *bookingService) GetAll() ([]booking.Core, error) {
	res, err := bs.qry.GetAll()
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return res, nil
}

// GetByUserID implements booking.Service.
func (bs *bookingService) GetByUserID(userID int) ([]booking.Core, error) {
	res, err := bs.qry.GetByUserID(userID)
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return res, nil
}
