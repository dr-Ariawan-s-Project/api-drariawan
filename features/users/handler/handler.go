package handler

import (
	"log"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
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
func (uh *UserHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestBody := UserRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		res, err := uh.srv.Insert(*ReqToCore(requestBody))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		log.Println(res)

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": "string"})
		return c.JSON(httpCode, mapResponse)

	}
}

// Update implements users.UserHandler.
func (uh *UserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestBody := UserRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		err = uh.srv.Update(*ReqToCore(requestBody), c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": "string"})
		return c.JSON(httpCode, mapResponse)
	}
}

// Delete implements users.UserHandler.
func (uh *UserHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := 1
		err := uh.srv.Delete(userID)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": "string"})
		return c.JSON(httpCode, mapResponse)
	}
}

// FindAll implements users.UserHandler.
func (uh *UserHandler) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uh.srv.FindAll("johndoe", 1, 1)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": CoreToResponseArray(res)})
		return c.JSON(httpCode, mapResponse)
	}
}

// FindById implements users.UserHandler.
func (uh *UserHandler) FindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uh.srv.FindById(1)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": CoreToResponse(res)})
		return c.JSON(httpCode, mapResponse)
	}
}

// FindByUsernameOrEmail implements users.UserHandler.
func (uh *UserHandler) FindByUsernameOrEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uh.srv.FindByUsernameOrEmail("johndoe")
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": CoreToResponse(res)})
		return c.JSON(httpCode, mapResponse)
	}
}
