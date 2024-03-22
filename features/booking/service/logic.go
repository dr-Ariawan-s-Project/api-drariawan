package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
)

type bookingService struct {
	qry booking.Data
	cfg *config.AppConfig
}

func New(sd booking.Data, cfg *config.AppConfig) booking.Service {
	return &bookingService{
		qry: sd,
		cfg: cfg,
	}
}

// for pagination
// GetPagination implements booking.Service.
func (bs *bookingService) GetPagination(page int, perPage int) (map[string]any, error) {
	totalRows, err := bs.qry.CountByFilter()
	response := map[string]any{
		"page":          0,
		"limit":         0,
		"total_pages":   0,
		"total_records": 0,
	}
	if err != nil {
		return response, err
	}
	paginationRes := helpers.CountPagination(totalRows, page, perPage)
	response["page"] = paginationRes.Page
	response["limit"] = paginationRes.Limit
	response["total_pages"] = paginationRes.TotalPages
	response["total_records"] = paginationRes.TotalRecords
	return response, nil
}

// Create implements booking.Service.
func (bs *bookingService) Create(data booking.Core, role string) error {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_PatientAccess {
		return errors.New(config.VAL_Unauthorized)
	}
	bookingID, err := bs.qry.Create(data)
	if err != nil {
		return errors.New(err.Error())
	}

	bookingResult, errBookingResult := bs.qry.GetByBookingID(*bookingID)
	if errBookingResult != nil {
		return errBookingResult
	}

	// format booking date
	layoutFormat := "2006-01-02T15:04:05+07:00"
	bookDate, _ := time.Parse(layoutFormat, bookingResult.BookingDate)
	y, m, d := bookDate.Date()
	bookDateStr := fmt.Sprintf("%d-%d-%d", d, m, y)

	// send data to func send email
	appointmentData := helpers.AppointmentDTO{
		BookingCode:       bookingResult.BookingCode,
		PatientName:       bookingResult.Patient.Name,
		Email:             bookingResult.Patient.Email,
		DoctorName:        bookingResult.Schedule.User.Name,
		Specialization:    bookingResult.Schedule.User.Specialization,
		HealthcareAddress: bookingResult.Schedule.HealthcareAddress,
		BookingDate:       bookDateStr,
		TimeStart:         bookingResult.Schedule.TimeStart,
		TimeEnd:           bookingResult.Schedule.TimeEnd,
	}

	// send email
	go helpers.SendMailAppointmentConfirmation(bookingResult.Patient.Email, bs.cfg.GMAIL_APP_PASSWORD, appointmentData)

	return nil
}

// Update implements schedule.bookingService.
func (bs *bookingService) Update(id string, data booking.Core, role string) error {
	err := bs.qry.Update(id, data)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete implements schedule.bookingService.
func (bs *bookingService) Delete(id string, role string) error {
	err := bs.qry.Delete(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// CancelBooking implements booking.Service.
func (bs *bookingService) CancelBooking(id string, role string) error {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess && strings.ToLower(role) != config.VAL_SuperAdminAccess && strings.ToLower(role) != config.VAL_AdminAccess {
		return errors.New(config.VAL_Unauthorized)
	}
	// check if bookingid is exist or not. err means not exist.
	_, errBookData := bs.qry.GetByBookingID(id)
	if errBookData != nil {
		return errBookData
	}

	// do cancel if bookingid exist
	errCancel := bs.qry.CancelBooking(id)
	return errCancel
}

// GetAll implements schedule.bookingService.
func (bs *bookingService) GetAll(role string, page int, perPage int) ([]booking.Core, error) {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess && strings.ToLower(role) != config.VAL_SuperAdminAccess && strings.ToLower(role) != config.VAL_AdminAccess {
		return []booking.Core{}, errors.New(config.VAL_Unauthorized)
	}
	if perPage <= 0 {
		perPage = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page * perPage) - perPage

	if offset < 0 {
		offset = 0
	}
	res, err := bs.qry.GetAll(offset, perPage)
	if err != nil {
		return []booking.Core{}, errors.New(err.Error())
	}
	return res, nil
}

// GetByUserID implements booking.Service.
func (bs *bookingService) GetByUserID(userID int, role string) ([]booking.Core, error) {
	if strings.ToLower(role) != config.VAL_SusterAccess && strings.ToLower(role) != config.VAL_DokterAccess && strings.ToLower(role) != config.VAL_PatientAccess && strings.ToLower(role) != config.VAL_SuperAdminAccess && strings.ToLower(role) != config.VAL_AdminAccess {
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
