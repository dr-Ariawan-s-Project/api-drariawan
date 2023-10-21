package handler

import (
	"strconv"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/schedule"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	echo "github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	srv schedule.ScheduleService
}

func New(sv schedule.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		srv: sv,
	}
}

// Create implements schedule.ScheduleHandler.
func (sh *ScheduleHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		requestBody := ScheduleRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_SCHEDULE_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		err = sh.srv.Create(*ReqToCore(requestBody))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_SCHEDULE_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] create data", config.FEAT_SCHEDULE_CODE, nil)
		return c.JSON(httpCode, mapResponse)

	}
}

// Update implements schedule.ScheduleHandler.
func (sh *ScheduleHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		scheID, _ := strconv.Atoi(c.QueryParam("id"))
		requestBody := ScheduleRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_SCHEDULE_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		err = sh.srv.Update(scheID, *ReqToCore(requestBody))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_SCHEDULE_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] update data", config.FEAT_SCHEDULE_CODE, nil)
		return c.JSON(httpCode, mapResponse)

	}
}

// Delete implements schedule.ScheduleHandler.
func (sh *ScheduleHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		strIdParam := c.QueryParam("id")
		scheID, _ := strconv.Atoi(strIdParam)
		err := sh.srv.Delete(scheID)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_SCHEDULE_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] delete data", config.FEAT_SCHEDULE_CODE, nil)
		return c.JSON(httpCode, mapResponse)
	}
}

// GetAll implements schedule.ScheduleHandler.
func (sh *ScheduleHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := sh.srv.GetAll()
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_SCHEDULE_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_SCHEDULE_CODE, res)
		return c.JSON(httpCode, mapResponse)

	}
}
