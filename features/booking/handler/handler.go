package handler

import (
	"strconv"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	echo "github.com/labstack/echo/v4"
)

type BookingHandler struct {
	srv booking.Service
}

func New(sv booking.Service) *BookingHandler {
	return &BookingHandler{
		srv: sv,
	}
}

// Create implements Booking.BookingHandler.
func (bh *BookingHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestBody := BookingRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		err = bh.srv.Create(*ReqToCore(requestBody))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": "string"})
		return c.JSON(httpCode, mapResponse)

	}
}

// Update implements Booking.BookingHandler.
func (bh *BookingHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		bookID, _ := strconv.Atoi(c.QueryParam("id"))
		requestBody := BookingRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		err = bh.srv.Update(bookID, *ReqToCore(requestBody))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": "string"})
		return c.JSON(httpCode, mapResponse)

	}
}

// Delete implements Booking.BookingHandler.
func (bh *BookingHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		strIdParam := c.QueryParam("id")
		bookID, _ := strconv.Atoi(strIdParam)
		err := bh.srv.Delete(bookID)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": "string"})
		return c.JSON(httpCode, mapResponse)
	}
}

// GetAll implements Booking.BookingHandler.
func (bh *BookingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bh.srv.GetAll()
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, map[string]interface{}{"data": res})
		return c.JSON(httpCode, mapResponse)

	}
}
