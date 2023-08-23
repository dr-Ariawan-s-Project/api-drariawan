package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_authData "github.com/dr-ariawan-s-project/api-drariawan/features/auth/data"
	_authHandler "github.com/dr-ariawan-s-project/api-drariawan/features/auth/handler"
	_authService "github.com/dr-ariawan-s-project/api-drariawan/features/auth/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_authHandler.AuthHandler {
	repo := _authData.New(db, cfg)
	service := _authService.New(repo)
	handler := _authHandler.New(service)

	return handler
}
