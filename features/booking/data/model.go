package data

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
)

type Bookings struct {
	ID          int
	PatientId   string
	ScheduleId  int
	BookingDate string
	State       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Patient     Patients
	Schedule    Schedules
}

type Patients struct {
	ID   string
	Name string
}

type Schedules struct {
	ID                int
	UserId            int
	HealthcareAddress string
	Day               string
	TimeStart         string
	TimeEnd           string
	User              Users
}
type Users struct {
	ID             int
	Name           string
	UrlPicture     string
	Specialization string
}

func CoreToData(core booking.Core) Bookings {
	return Bookings{
		ID:          core.ID,
		PatientId:   core.PatientId,
		ScheduleId:  core.ScheduleId,
		BookingDate: core.BookingDate,
		State:       core.State,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
		DeletedAt:   core.DeletedAt,
		Patient: Patients{
			ID:   core.Patient.ID,
			Name: core.Patient.Name,
		},
		Schedule: Schedules{
			ID:                core.Schedule.ID,
			UserId:            core.Schedule.UserId,
			HealthcareAddress: core.Schedule.HealthcareAddress,
			Day:               core.Schedule.Day,
			TimeStart:         core.Schedule.TimeStart,
			TimeEnd:           core.Schedule.TimeEnd,
			User: Users{
				ID:             core.Schedule.User.ID,
				Name:           core.Schedule.User.Name,
				UrlPicture:     core.Schedule.User.UrlPicture,
				Specialization: core.Schedule.User.Specialization,
			},
		},
	}
}

func DataToCore(data Bookings) booking.Core {
	return booking.Core{
		ID:          data.ID,
		PatientId:   data.PatientId,
		ScheduleId:  data.ScheduleId,
		BookingDate: data.BookingDate,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
		Patient: booking.Patients{
			ID:   data.Patient.ID,
			Name: data.Patient.Name,
		},
		Schedule: booking.Schedules{
			ID:                data.Schedule.ID,
			UserId:            data.Schedule.UserId,
			HealthcareAddress: data.Schedule.HealthcareAddress,
			Day:               data.Schedule.Day,
			TimeStart:         data.Schedule.TimeStart,
			TimeEnd:           data.Schedule.TimeEnd,
			User: booking.Users{
				ID:             data.Schedule.User.ID,
				Name:           data.Schedule.User.Name,
				UrlPicture:     data.Schedule.User.UrlPicture,
				Specialization: data.Schedule.User.Specialization,
			},
		},
	}
}

func DataToCoreArray(data []Bookings) []booking.Core {
	result := []booking.Core{}
	for _, val := range data {
		result = append(result, DataToCore(val))
	}
	return result
}
