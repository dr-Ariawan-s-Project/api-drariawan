package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/booking"

type BookingRequest struct {
	PatientId   string `json:"patient_id" form:"patient_id"`
	ScheduleId  int    `json:"schedule_id" form:"schedule_id"`
	BookingDate string `json:"booking_date" form:"booking_date"`
}

func ReqToCore(data interface{}) *booking.Core {
	res := booking.Core{}
	switch data.(type) {
	case BookingRequest:
		cnv := data.(BookingRequest)
		res.PatientId = cnv.PatientId
		res.ScheduleId = cnv.ScheduleId
		res.BookingDate = cnv.BookingDate

	default:
		return nil
	}

	return &res
}
