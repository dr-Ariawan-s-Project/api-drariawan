package handler

import (
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
)

type PatientResponse struct {
	ID             string           `json:"id"`
	Name           string           `json:"name,omitempty"`
	Email          string           `json:"email,omitempty"`
	NIK            string           `json:"nik,omitempty"`
	DOB            *string          `json:"dob,omitempty"`
	Phone          string           `json:"phone,omitempty"`
	Gender         *string          `json:"gender,omitempty"`
	MarriageStatus string           `json:"marriage_status,omitempty"`
	Nationality    string           `json:"nationality,omitempty"`
	PartnerID      *string          `json:"partner_id,omitempty"`
	Partner        *PatientResponse `json:"partner,omitempty"`
}

func CoreToResponse(dataCore patient.Core) PatientResponse {
	response := PatientResponse{
		ID:             dataCore.ID,
		Name:           dataCore.Name,
		Email:          dataCore.Email,
		NIK:            dataCore.NIK,
		Phone:          dataCore.Phone,
		Gender:         dataCore.Gender,
		MarriageStatus: dataCore.MarriageStatus,
		Nationality:    dataCore.Nationality,
		PartnerID:      dataCore.PartnerID,
	}

	if dataCore.DOB != nil {
		dob := dataCore.DOB.Format("2006-01-02")
		response.DOB = &dob
	}

	if dataCore.Partner != nil {
		dataPartner := CoreToResponse(*dataCore.Partner)
		response.Partner = &dataPartner
	}

	return response
}

func CoreToResponseList(dataCore []patient.Core) []PatientResponse {
	var result []PatientResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
