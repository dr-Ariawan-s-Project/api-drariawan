package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_bookingData "github.com/dr-ariawan-s-project/api-drariawan/features/booking/data"
	_bookingHandler "github.com/dr-ariawan-s-project/api-drariawan/features/booking/handler"
	_bookingService "github.com/dr-ariawan-s-project/api-drariawan/features/booking/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_bookingHandler.BookingHandler {
	repo := _bookingData.New(db)
	service := _bookingService.New(repo, cfg)
	handler := _bookingHandler.New(service)

	return handler
}
