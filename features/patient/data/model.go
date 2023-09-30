package data

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"gorm.io/gorm"
)

type Patient struct {
	ID             string
	Name           string
	Email          string
	Password       string
	NIK            string
	DOB            *time.Time
	Phone          string
	Gender         *string
	MarriageStatus string
	Nationality    string
	PartnerID      *string
	Partner        *Patient `gorm:"foreignKey:PartnerID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func PatientCoreToModel(data patient.Core) Patient {
	return Patient{
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		NIK:            data.NIK,
		DOB:            data.DOB,
		Phone:          data.Phone,
		Gender:         data.Gender,
		MarriageStatus: data.MarriageStatus,
		Nationality:    data.Nationality,
	}
}

func (data Patient) ModelToCore() patient.Core {
	response := patient.Core{
		ID:             data.ID,
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		NIK:            data.NIK,
		DOB:            data.DOB,
		Phone:          data.Phone,
		Gender:         data.Gender,
		MarriageStatus: data.MarriageStatus,
		Nationality:    data.Nationality,
		PartnerID:      data.PartnerID,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
	if data.Partner != nil {
		response.Partner = &patient.Core{
			ID:   data.Partner.ID,
			Name: data.Partner.Name,
		}
	}
	return response
}

func ListModelToCore(data []Patient) []patient.Core {
	var dataCore []patient.Core
	for _, v := range data {
		dataCore = append(dataCore, v.ModelToCore())
	}
	return dataCore
}
