package factory

import (
	_usersData "github.com/dr-ariawan-s-project/api-drariawan/features/users/data"
	_usersHandler "github.com/dr-ariawan-s-project/api-drariawan/features/users/handler"
	_usersService "github.com/dr-ariawan-s-project/api-drariawan/features/users/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *_usersHandler.UserHandler {
	repo := _usersData.New(db)
	service := _usersService.New(repo)
	handler := _usersHandler.New(service)

	return handler
}
