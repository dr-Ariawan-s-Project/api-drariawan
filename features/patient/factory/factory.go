package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	_patientData "github.com/dr-ariawan-s-project/api-drariawan/features/patient/data"
	_patientHandler "github.com/dr-ariawan-s-project/api-drariawan/features/patient/handler"
	_patientService "github.com/dr-ariawan-s-project/api-drariawan/features/patient/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_patientHandler.PatientHandler {
	repo := _patientData.New(db)
	service := _patientService.New(repo, cfg)
	handler := _patientHandler.New(service)
	return handler
}

func NewServiceFactory(db *gorm.DB, cfg *config.AppConfig) *patient.PatientServiceInterface {
	repo := _patientData.New(db)
	service := _patientService.New(repo, cfg)
	return &service
}
