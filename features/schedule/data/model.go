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
		DeletedAt:         &core.UpdatedAt,
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
		DeletedAt:         &data.UpdatedAt,
	}
}

func DataToCoreArray(data []Schedules) []schedule.Core {
	result := []schedule.Core{}
	for _, val := range data {
		result = append(result, DataToCore(val))
	}
	return result
}
