package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_dashboardHandler "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard/handler"
	_dashboardService "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard/service"
	_patientFactory "github.com/dr-ariawan-s-project/api-drariawan/features/patient/factory"
	_questionaireFactory "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/factory"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_dashboardHandler.DashboardHandler {
	patientService := _patientFactory.NewServiceFactory(db, cfg)
	questionerService := _questionaireFactory.NewServiceFactory(db, cfg)

	service := _dashboardService.New(*questionerService, *patientService)
	handler := _dashboardHandler.New(service)

	return handler
}
