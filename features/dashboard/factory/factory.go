package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_dashboardData "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard/data"
	_dashboardHandler "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard/handler"
	_dashboardService "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_dashboardHandler.DashboardHandler {

	data := _dashboardData.New(db)
	service := _dashboardService.New(data)
	handler := _dashboardHandler.New(service)

	return handler
}
