package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/patient"

type PatientRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	NIK            string `json:"nik"`
	DOB            string `json:"dob"`
	Phone          string `json:"phone"`
	Gender         string `json:"gender"`
	MarriageStatus string `json:"marriage_status"` // male or female
	Nationality    string `json:"nationality"`
	PartnerEmail   string `json:"partner_email"`
}

func (data PatientRequest) RequestToCore() patient.Core {
	return patient.Core{
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		NIK:            data.NIK,
		Phone:          data.Phone,
		MarriageStatus: data.MarriageStatus,
		Nationality:    data.Nationality,
	}
}
