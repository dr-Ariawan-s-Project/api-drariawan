package handler

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/app/middlewares"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/labstack/echo/v4"
)

type PatientHandler struct {
	patientService patient.PatientServiceInterface
}

func New(srv patient.PatientServiceInterface) *PatientHandler {
	return &PatientHandler{
		patientService: srv,
	}
}

func (handler *PatientHandler) AddPatient(c echo.Context) error {
	input := new(PatientRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		log.Println("error bind", errBind)
		jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.REQ_ErrorBindData), config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	var dataCore = input.RequestToCore()
	if input.DOB != "" {
		layoutFormat := "2006-01-02"
		dobTime, errParse := time.Parse(layoutFormat, input.DOB)
		if errParse != nil {
			log.Println("error format dob", errParse)
			jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.REQ_InvalidParam), config.FEAT_PATIENT_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		dataCore.DOB = &dobTime
	}

	if input.Gender != "" {
		dataCore.Gender = &input.Gender
	}

	dataNewPatient, err := handler.patientService.Insert(dataCore, input.PartnerEmail)
	if err != nil {
		log.Println("error", err)
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	dataResponse := CoreToResponse(*dataNewPatient)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] add patient", config.FEAT_PATIENT_CODE, dataResponse)
	return c.JSON(httpCode, mapResponse)
}

func (handler *PatientHandler) GetAll(c echo.Context) error {
	qSearch := c.QueryParam("search")
	qPage := c.QueryParam("page")
	qLimit := c.QueryParam("limit")

	var page, limit int
	if qPage != "" {
		var errPage error
		page, errPage = strconv.Atoi(qPage)
		if errPage != nil {
			jsonResponse, httpCode := helpers.WebResponseError(errPage, config.FEAT_PATIENT_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
	}
	if qLimit != "" {
		var errLimit error
		limit, errLimit = strconv.Atoi(qLimit)
		if errLimit != nil {
			jsonResponse, httpCode := helpers.WebResponseError(errLimit, config.FEAT_PATIENT_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
	}
	results, err := handler.patientService.FindAll(qSearch, page, limit)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	var patientRespose = CoreToResponseList(results)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_PATIENT_CODE, patientRespose)
	return c.JSON(httpCode, mapResponse)
}

func (handler *PatientHandler) GetById(c echo.Context) error {
	id := c.Param("patient_id")
	result, err := handler.patientService.FindById(id)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	var patientRespose = CoreToResponse(*result)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_PATIENT_CODE, patientRespose)
	return c.JSON(httpCode, mapResponse)
}

func (handler *PatientHandler) EditPatient(c echo.Context) error {
	id := c.Param("patient_id")
	input := new(PatientRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errBind, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	var dataCore = input.RequestToCore()
	dataCore.ID = id
	if input.DOB != "" {
		layoutFormat := "2006-01-02"
		dobTime, errParse := time.Parse(layoutFormat, input.DOB)
		if errParse != nil {
			jsonResponse, httpCode := helpers.WebResponseError(errParse, config.FEAT_PATIENT_CODE)
			return c.JSON(httpCode, jsonResponse)
		}
		dataCore.DOB = &dobTime
	}

	if input.Gender != "" {
		dataCore.Gender = &input.Gender
	}
	dataPatient, err := handler.patientService.Update(dataCore)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	var patientRespose = CoreToResponse(*dataPatient)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] update data", config.FEAT_PATIENT_CODE, patientRespose)
	return c.JSON(httpCode, mapResponse)
}

func (handler *PatientHandler) DeleteById(c echo.Context) error {
	id := c.Param("patient_id")
	err := handler.patientService.Delete(id)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] delete data", config.FEAT_PATIENT_CODE, nil)
	return c.JSON(httpCode, mapResponse)
}

func (handler *PatientHandler) GetProfile(c echo.Context) error {
	idToken, roleToken, errToken := middlewares.ExtractTokenJWT(c)
	if errToken != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.JWT_FailedCastingJwtToken), config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	if roleToken != config.VAL_PatientAccess {
		jsonResponse, httpCode := helpers.WebResponseError(errors.New(config.VAL_Unauthorized), config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	id := middlewares.ConvertPatientID(idToken)
	result, err := handler.patientService.FindById(id)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_PATIENT_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	var patientRespose = CoreToResponse(*result)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_PATIENT_CODE, patientRespose)
	return c.JSON(httpCode, mapResponse)
}
