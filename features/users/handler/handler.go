package handler

import (
	"errors"
	"log"
	"strconv"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/app/middlewares"
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
		_, role, errToken := middlewares.ExtractTokenJWT(c)
		log.Println("role:", role)
		if errToken != nil {
			jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.JWT_FailedCastingJwtToken), config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		requestBody := UserRequest{}
		err := c.Bind(&requestBody)
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

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] create data", config.FEAT_USER_CODE, nil)
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
		IDUser := middlewares.ConvertUserID(UserID)
		requestBody := UserRequest{}
		err = c.Bind(&requestBody)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		err = uh.srv.Update(*ReqToCore(requestBody), IDUser)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		mapResponse, httpCode := helpers.WebResponseSuccess("[success] update data", config.FEAT_USER_CODE, nil)
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
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] delete data", config.FEAT_USER_CODE, nil)
		return c.JSON(httpCode, mapResponse)
	}
}

// FindAll implements users.UserHandler.
func (uh *UserHandler) FindAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("search")
		rp := c.QueryParam("rp")
		page := c.QueryParam("page")

		//check query param
		var pageInt, rpInt int
		if page != "" {
			var errPage error
			pageInt, errPage = strconv.Atoi(page)
			if errPage != nil {
				jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.REQ_InvalidPageParam), config.FEAT_PATIENT_CODE)
				return c.JSON(httpCode, jsonResponse)
			}
		}
		if rp != "" {
			var errLimit error
			rpInt, errLimit = strconv.Atoi(rp)
			if errLimit != nil {
				jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.REQ_InvalidLimitParam), config.FEAT_PATIENT_CODE)
				return c.JSON(httpCode, jsonResponse)
			}
		}

		res, err := uh.srv.FindAll(search, rpInt, pageInt)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, CoreToResponseArray(res))

		//pagination
		paginationRes, errPagination := uh.srv.GetPagination(search, pageInt, rpInt)
		if errPagination == nil {
			mapResponse["pagination"] = paginationRes
		}
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
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, CoreToResponse(res))
		return c.JSON(httpCode, mapResponse)
	}
}

// FindById implements users.UserHandler.
func (uh *UserHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID, _, err := encrypt.ExtractToken(c.Get("user"))
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		IDUser := middlewares.ConvertUserID(UserID)
		res, err := uh.srv.FindById(IDUser)
		if err != nil {
			jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_USER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_USER_CODE, CoreToResponse(res))
		return c.JSON(httpCode, mapResponse)
	}
}
