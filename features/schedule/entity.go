package schedule

import (
	"time"
)

type Core struct {
	ID                int        `json:"schedule_id"`
	UserId            int        `json:"user_id"`
	HealthcareAddress string     `json:"health_care_address"`
	Day               string     `json:"day"`
	TimeStart         string     `json:"time_start"`
	TimeEnd           string     `json:"time_end"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	User              User       `json:"user"`
	Booking           []Booking  `json:"booking"`
}

type Booking struct {
	ID          string     `json:"id"`
	BookingCode string     `json:"booking_code"`
	PatientId   string     `json:"patient_id"`
	ScheduleId  int        `json:"schedule_id"`
	BookingDate string     `json:"booking_date"`
	State       string     `json:"state"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type User struct {
	ID             int    `json:"user_id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	UrlPicture     string `json:"picture"`
	Specialization string `json:"specialization"`
}

type ScheduleService interface {
	Create(data Core, role string) error
	Update(id int, data Core, role string) error
	Delete(id int, role string) error
	GetAll(role string, page int, perPage int) ([]Core, error)
	GetPagination(page int, perPage int) (map[string]any, error)
}
type ScheduleData interface {
	Create(data Core) error
	Update(id int, data Core) error
	Delete(id int) error
	GetAll(offset int, limit int) ([]Core, error)
	CountByFilter() (int64, error)
}
