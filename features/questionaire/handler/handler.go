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
