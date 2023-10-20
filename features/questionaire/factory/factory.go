package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_patientFactory "github.com/dr-ariawan-s-project/api-drariawan/features/patient/factory"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	_questionaireData "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/data"
	_questionaireHandler "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"
	_questionaireService "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_questionaireHandler.QuestionaireHandler {
	patientService := _patientFactory.NewServiceFactory(db, cfg)

	repo := _questionaireData.New(db)
	service := _questionaireService.New(repo, *patientService, cfg)
	handler := _questionaireHandler.New(service)

	return handler
}

func NewServiceFactory(db *gorm.DB, cfg *config.AppConfig) *questionaire.QuestionaireServiceInterface {
	patientService := _patientFactory.NewServiceFactory(db, cfg)
	repo := _questionaireData.New(db)
	service := _questionaireService.New(repo, *patientService, cfg)
	return &service
}
