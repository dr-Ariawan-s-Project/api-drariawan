package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/schedule"

type ScheduleRequest struct {
	UserId            int    `json:"user_id" form:"user_id"`
	HealthcareAddress string `json:"health_care_address" form:"health_care_address"`
	Day               string `json:"day" form:"day"`
	TimeStart         string `json:"time_start" form:"time_start"`
	TimeEnd           string `json:"time_end" form:"time_end"`
}

func ReqToCore(data interface{}) *schedule.Core {
	res := schedule.Core{}
	switch data.(type) {
	case ScheduleRequest:
		cnv := data.(ScheduleRequest)
		res.UserId = cnv.UserId
		res.HealthcareAddress = cnv.HealthcareAddress
		res.Day = cnv.Day
		res.TimeStart = cnv.TimeStart
		res.TimeEnd = cnv.TimeEnd

	default:
		return nil
	}

	return &res
}
