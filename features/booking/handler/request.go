package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/booking"

type BookingRequest struct {
	PatientId  string `json:"patient_id" form:"patient_id"`
	ScheduleId int    `json:"schedule_id" form:"schedule_id"`
}

func ReqToCore(data interface{}) *booking.Core {
	res := booking.Core{}
	switch data.(type) {
	case BookingRequest:
		cnv := data.(BookingRequest)
		res.PatientId = cnv.PatientId
		res.ScheduleId = cnv.ScheduleId

	default:
		return nil
	}

	return &res
}
