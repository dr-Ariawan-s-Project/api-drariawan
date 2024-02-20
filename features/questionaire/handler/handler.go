package handler

import (
	"strconv"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/labstack/echo/v4"
)

type QuestionaireHandler struct {
	questionaireService questionaire.QuestionaireServiceInterface
}

func New(service questionaire.QuestionaireServiceInterface) *QuestionaireHandler {
	return &QuestionaireHandler{
		questionaireService: service,
	}
}

func (handler *QuestionaireHandler) Validate(c echo.Context) error {
	validateInput := new(ValidateRequest)
	errBind := c.Bind(&validateInput)
	if errBind != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errBind, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	patientData := questionaire.Patient{
		Email: validateInput.Email,
		Phone: validateInput.Phone,
	}

	codeAttempt, countAttempt, err := handler.questionaireService.Validate(patientData, validateInput.As, validateInput.PartnerEmail)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		if countAttempt != 0 {
			jsonResponse["messages"] = []string{"user has already taken the test"}
		}
		return c.JSON(httpCode, jsonResponse)
	}
	data := map[string]any{
		"code_attempt":  codeAttempt,
		"count_attempt": countAttempt,
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] test attempt added. Start your test.", config.FEAT_QUESTIONAIRE_CODE, data)
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) CheckIsValidCodeAttempt(c echo.Context) error {
	codeAttempt := c.QueryParam("code")
	isValid, err := handler.questionaireService.CheckIsValidCodeAttempt(codeAttempt)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	data := map[string]any{}
	if isValid {
		data["is_valid"] = true
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] code is valid", config.FEAT_QUESTIONAIRE_CODE, data)
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) AddAnswer(c echo.Context) error {
	answerInput := new(AnswerRequest)
	errBind := c.Bind(&answerInput)
	if errBind != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errBind, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	err := handler.questionaireService.InsertAnswer(answerInput.CodeAttempt, answerInput.RequestToCore())
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] add answer", config.FEAT_QUESTIONAIRE_CODE, nil)
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) GetAllTestAttempt(c echo.Context) error {
	qStatus := c.QueryParam("status")
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
	results, err := handler.questionaireService.GetTestAttempt(qStatus, page, limit)

	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	var attemptResponse = ListAttemptCoreToResponse(results)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_QUESTIONAIRE_CODE, attemptResponse)
	//pagination
	paginationRes, errPagination := handler.questionaireService.GetPaginationTestAttempt(qStatus, page, limit)
	if errPagination == nil {
		mapResponse["pagination"] = paginationRes
	}
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) GetTestAttemptById(c echo.Context) error {
	id := c.Param("attempt_id")
	result, err := handler.questionaireService.GetTestAttemptById(id)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	var attemptResponse = AttemptCoreToResponse(*result)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_QUESTIONAIRE_CODE, attemptResponse)
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) GetAllAnswerByAttempt(c echo.Context) error {
	idAttempt := c.Param("attempt_id")

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
	results, err := handler.questionaireService.GetAllAnswerByAttempt(idAttempt, page, limit)

	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	var answerResponse = ListAnswerCoreToResponse(results)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_QUESTIONAIRE_CODE, answerResponse)
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) AddAssesmentByAttempt(c echo.Context) error {
	idAttempt := c.Param("attempt_id")
	assesmentInput := new(AssesmentRequest)
	errBind := c.Bind(&assesmentInput)
	if errBind != nil {
		jsonResponse, httpCode := helpers.WebResponseError(errBind, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	data := assesmentInput.RequestToCore()
	data.Id = idAttempt

	err := handler.questionaireService.InsertAssesment(data)
	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] add assesment", config.FEAT_QUESTIONAIRE_CODE, nil)
	return c.JSON(httpCode, mapResponse)
}

func (handler *QuestionaireHandler) GetAllQuestion(c echo.Context) error {
	results, err := handler.questionaireService.GetAll()

	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	var questionRespose = CoreToResponseList(results)
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_QUESTIONAIRE_CODE, questionRespose)
	return c.JSON(httpCode, mapResponse)
}
