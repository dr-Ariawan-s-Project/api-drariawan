package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_scheduleData "github.com/dr-ariawan-s-project/api-drariawan/features/schedule/data"
	_scheduleHandler "github.com/dr-ariawan-s-project/api-drariawan/features/schedule/handler"
	_scheduleService "github.com/dr-ariawan-s-project/api-drariawan/features/schedule/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_scheduleHandler.ScheduleHandler {
	repo := _scheduleData.New(db)
	service := _scheduleService.New(repo)
	handler := _scheduleHandler.New(service)

	return handler
}
