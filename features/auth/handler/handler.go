package handler

import (
	"net/http"
	"reflect"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/auth"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/validation"
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

func (handler *AuthHandler) Login(c echo.Context) error {
	authInput := new(loginRequest)
	errBind := c.Bind(&authInput)
	if errBind != nil {
		errCode, layerCode, errMessage := helpers.CheckHandlerErrorCode(errBind)
		code := helpers.GenerateCodeResponse(errCode, config.FEAT_AUTH_CODE, layerCode)
		return c.JSON(errCode, helpers.WebResponseError(errMessage.Error(), code))
	}

	token, err := handler.authService.Login(authInput.Email, authInput.Password)
	if err != nil {
		errCode, layerCode, errMessage := helpers.CheckHandlerErrorCode(err)
		code := helpers.GenerateCodeResponse(errCode, config.FEAT_AUTH_CODE, layerCode)
		//if error from validation
		if reflect.TypeOf(err).String() == "validation.ValidationError" {
			valErr := err.(validation.ValidationError)
			return c.JSON(errCode, helpers.WebResponseError(valErr.Errors, code))
		}

		return c.JSON(errCode, helpers.WebResponseError(errMessage.Error(), code))
	}

	code := helpers.GenerateCodeResponse(http.StatusOK, config.FEAT_AUTH_CODE, config.RESPONSE_SUCCESS_CODE)
	return c.JSON(http.StatusOK, helpers.WebResponseSuccess("[success] login", code, map[string]any{
		"token": token,
	}))

}
