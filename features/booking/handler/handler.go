package handler

import (
	"errors"
	"strconv"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/booking"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
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
		_, role, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		requestBody := BookingRequest{}
		err = c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		err = bh.srv.Create(*ReqToCore(requestBody), role)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] create data", config.FEAT_BOOKING_CODE, nil)
		return c.JSON(httpCode, mapResponse)

	}
}

// Update implements Booking.BookingHandler.
func (bh *BookingHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		bookID := c.Param("bookingid")
		requestBody := BookingRequest{}
		err = c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		err = bh.srv.Update(bookID, *ReqToCore(requestBody), role)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] update data", config.FEAT_BOOKING_CODE, nil)
		return c.JSON(httpCode, mapResponse)

	}
}

// Delete implements Booking.BookingHandler.
func (bh *BookingHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		bookID := c.Param("bookingid")
		err = bh.srv.Delete(bookID, role)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] delete data", config.FEAT_BOOKING_CODE, nil)
		return c.JSON(httpCode, mapResponse)
	}
}

// GetAll implements Booking.BookingHandler.
func (bh *BookingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		qPage := c.QueryParam("page")
		qLimit := c.QueryParam("limit")

		var page, limit int
		if qPage != "" {
			var errPage error
			page, errPage = strconv.Atoi(qPage)
			if errPage != nil {
				jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.REQ_InvalidPageParam), config.FEAT_PATIENT_CODE)
				return c.JSON(httpCode, jsonResponse)
			}
		}
		if qLimit != "" {
			var errLimit error
			limit, errLimit = strconv.Atoi(qLimit)
			if errLimit != nil {
				jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.REQ_InvalidLimitParam), config.FEAT_PATIENT_CODE)
				return c.JSON(httpCode, jsonResponse)
			}
		}

		res, err := bh.srv.GetAll(role, page, limit)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_BOOKING_CODE, res)
		//pagination
		paginationRes, errPagination := bh.srv.GetPagination(page, limit)
		if errPagination == nil {
			mapResponse["pagination"] = paginationRes
		}
		return c.JSON(httpCode, mapResponse)

	}
}

func (bh *BookingHandler) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		strIdParam := c.Param("userid")
		userID, _ := strconv.Atoi(strIdParam)
		res, err := bh.srv.GetByUserID(userID, role)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_BOOKING_CODE, res)
		return c.JSON(httpCode, mapResponse)

	}
}

func (bh *BookingHandler) GetByPatientID() echo.HandlerFunc {
	return func(c echo.Context) error {
		strIdParam := c.QueryParam("patient_id")
		res, err := bh.srv.GetByPatientID(strIdParam)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_BOOKING_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_BOOKING_CODE, res)
		return c.JSON(httpCode, mapResponse)

	}
}
