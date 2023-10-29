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
	User              User
}

type User struct {
	ID             int
	Name           string
	Email          string
	Phone          string
	UrlPicture     string
	Specialization string
}

type ScheduleService interface {
	Create(data Core, role string) error
	Update(id int, data Core, role string) error
	Delete(id int, role string) error
	GetAll(role string) ([]Core, error)
}
type ScheduleData interface {
	Create(data Core) error
	Update(id int, data Core) error
	Delete(id int) error
	GetAll() ([]Core, error)
}
