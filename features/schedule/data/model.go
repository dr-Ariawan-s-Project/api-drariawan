package data

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
)

type Schedules struct {
	ID                int
	UserId            int
	HealthcareAddress string
	Day               string
	TimeStart         string
	TimeEnd           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	User              Users
	Booking           []Booking `gorm:"foreignKey:ScheduleId"`
}

type Booking struct {
	ID          string
	BookingCode string
	PatientId   string
	ScheduleId  int
	BookingDate string
	State       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type Users struct {
	ID             int
	Name           string
	Email          string
	Phone          string
	UrlPicture     string
	Specialization string
}

func BookDataToBookCore(data Booking) schedule.Booking {
	return schedule.Booking{
		ID:          data.ID,
		BookingCode: data.BookingCode,
		PatientId:   data.PatientId,
		ScheduleId:  data.ScheduleId,
		BookingDate: data.BookingDate,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
	}
}

func CoreToData(core schedule.Core) Schedules {
	return Schedules{
		ID:                core.ID,
		UserId:            core.UserId,
		HealthcareAddress: core.HealthcareAddress,
		Day:               core.Day,
		TimeStart:         core.TimeStart,
		TimeEnd:           core.TimeEnd,
		CreatedAt:         core.CreatedAt,
		UpdatedAt:         core.UpdatedAt,
		DeletedAt:         core.DeletedAt,
		User: Users{
			ID:             core.User.ID,
			Name:           core.User.Name,
			Email:          core.User.Email,
			Phone:          core.User.Phone,
			UrlPicture:     core.User.UrlPicture,
			Specialization: core.User.Specialization,
		},
	}
}

func DataToCore(data Schedules) schedule.Core {
	return schedule.Core{
		ID:                data.ID,
		UserId:            data.UserId,
		HealthcareAddress: data.HealthcareAddress,
		Day:               data.Day,
		TimeStart:         data.TimeStart,
		TimeEnd:           data.TimeEnd,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
		DeletedAt:         data.DeletedAt,
		User: schedule.User{
			ID:             data.User.ID,
			Name:           data.User.Name,
			Email:          data.User.Email,
			Phone:          data.User.Phone,
			UrlPicture:     data.User.UrlPicture,
			Specialization: data.User.Specialization,
		},
		Booking: BookingDataToCoreArray(data.Booking),
	}
}

func DataToCoreArray(data []Schedules) []schedule.Core {
	result := []schedule.Core{}
	for _, val := range data {
		result = append(result, DataToCore(val))
	}
	return result
}

func BookingDataToCoreArray(data []Booking) []schedule.Booking {
	result := []schedule.Booking{}
	for _, val := range data {
		result = append(result, BookDataToBookCore(val))
	}
	return result
}
