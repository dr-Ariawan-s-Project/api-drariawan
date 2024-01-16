package booking

import (
	"time"
)

type Core struct {
	ID          string     `json:"id"`
	BookingCode string     `json:"booking_code"`
	PatientId   string     `json:"patient_id"`
	ScheduleId  int        `json:"schedule_id"`
	BookingDate string     `json:"booking_date"`
	State       string     `json:"state"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Patient     Patients   `json:"patient"`
	Schedule    Schedules  `json:"schedule"`
}

type Patients struct {
	ID    string `json:"patient_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Schedules struct {
	ID                int    `json:"schedule_id"`
	UserId            int    `json:"user_id"`
	HealthcareAddress string `json:"health_care_address"`
	Day               string `json:"day"`
	TimeStart         string `json:"time_start"`
	TimeEnd           string `json:"time_end"`
	User              Users  `json:"user"`
}
type Users struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	UrlPicture     string `json:"picture"`
	Specialization string `json:"specialization"`
}

type Service interface {
	Create(data Core, role string) error
	Update(id string, data Core, role string) error
	Delete(id string, role string) error
	GetAll(role string, page int, perPage int) ([]Core, error)
	GetByUserID(userID int, role string) ([]Core, error)
	GetByPatientID(patientID string) ([]Core, error)
	GetPagination(page int, perPage int) (map[string]any, error)
}
type Data interface {
	Create(data Core) (bookingID *string, err error)
	Update(id string, data Core) error
	Delete(id string) error
	GetAll(offset int, limit int) ([]Core, error)
	GetByUserID(userID int) ([]Core, error)
	GetByPatientID(patientID string) ([]Core, error)
	GetByBookingID(bookingID string) (*Core, error)
	CountByFilter() (int64, error)
}
