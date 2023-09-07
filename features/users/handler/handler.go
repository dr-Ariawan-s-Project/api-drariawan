package handler

import (
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UserService
}

func New(us users.UserService) *UserHandler {
	return &UserHandler{
		srv: us,
	}
}

// Insert implements users.UserHandler.
func (*UserHandler) Insert() echo.HandlerFunc {
	panic("unimplemented")
}

// Update implements users.UserHandler.
func (*UserHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements users.UserHandler.
func (*UserHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}
