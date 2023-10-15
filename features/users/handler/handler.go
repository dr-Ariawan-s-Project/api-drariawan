package handler

import (
	"log"
	"strconv"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/users"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	echo "github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.Service
}

func New(us users.Service) *UserHandler {
	return &UserHandler{
		srv: us,
	}
}

// Insert implements users.UserHandler.
func (uh *UserHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		requestBody := UserRequest{}
		err = c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		res, err := uh.srv.Insert(*ReqToCore(requestBody), role)
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
		UserID, _, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		requestBody := UserRequest{}
		err = c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		err = uh.srv.Update(*ReqToCore(requestBody), UserID)
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
		strIdParam := c.QueryParam("id")
		userID, _ := strconv.Atoi(strIdParam)
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
		search := c.QueryParam("search")
		rp, _ := strconv.Atoi(c.QueryParam("rp"))
		page, _ := strconv.Atoi(c.QueryParam("page"))

		res, err := uh.srv.FindAll(search, rp, page)
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
		userID, _ := strconv.Atoi(c.QueryParam("id"))
		res, err := uh.srv.FindById(userID)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": CoreToResponse(res)})
		return c.JSON(httpCode, mapResponse)
	}
}

// FindById implements users.UserHandler.
func (uh *UserHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		res, err := uh.srv.FindById(userID)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": CoreToResponse(res)})
		return c.JSON(httpCode, mapResponse)
	}
}
