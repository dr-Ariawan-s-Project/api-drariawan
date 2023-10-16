package handler

import (
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

func (handler *QuestionaireHandler) GetDashboardQuestioner(c echo.Context) error {
	results, err := handler.questionaireService.GetDashboard()

	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_QUESTIONAIRE_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	var questionerDashboardRespose = DashboardQuestionerResponse{
		AllQuestioner:   results.AllQuestioner,
		NeedAssess:      results.NeedAssessQuestioner,
		MonthQuestioner: results.MonthQuestioner,
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_QUESTIONAIRE_CODE, questionerDashboardRespose)
	return c.JSON(httpCode, mapResponse)
}
