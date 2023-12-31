package handler

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/auth"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService auth.AuthServiceInterface
}

func New(service auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (handler *AuthHandler) LoginPatient(c echo.Context) error {
	authInput := new(loginRequest)
	errBind := c.Bind(&authInput)
	if errBind != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errBind, config.FEAT_AUTH_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	dataPatient, token, err := handler.authService.LoginPatient(authInput.Email, authInput.Password)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_AUTH_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	mapResponse, httpCode := helpers.WebResponseSuccess("[success] login", config.FEAT_AUTH_CODE, map[string]any{
		"token": token,
		"name":  &dataPatient.Name,
		"role":  config.VAL_PatientAccess,
	})
	return c.JSON(httpCode, mapResponse)

}

func (handler *AuthHandler) Login(c echo.Context) error {
	authInput := new(loginRequest)
	errBind := c.Bind(&authInput)
	if errBind != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errBind, config.FEAT_AUTH_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	dataUser, token, err := handler.authService.Login(authInput.Email, authInput.Password)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_AUTH_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	mapResponse, httpCode := helpers.WebResponseSuccess("[success] login", config.FEAT_AUTH_CODE, map[string]any{
		"token": token,
		"name":  &dataUser.Name,
		"role":  &dataUser.Role,
	})
	return c.JSON(httpCode, mapResponse)

}
