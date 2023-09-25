package schedule

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                int
	UserId            int
	HealthcareAddress string
	Day               string
	TimeStart         time.Time
	TimeEnd           time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

type ScheduleHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type ScheduleService interface {
	Create(data Core) error
	Update(id int, data Core) error
	Delete(id int) error
	GetAll() ([]Core, error)
}
type ScheduleData interface {
	Create(data Core) error
	Update(id int, data Core) error
	Delete(id int) error
	GetAll() ([]Core, error)
}
