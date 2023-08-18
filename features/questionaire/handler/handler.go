package handler

import (
	"net/http"

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
		errCode, layerCode, errMessage := helpers.CheckHandlerErrorCode(err)
		code := helpers.GenerateCodeResponse(errCode, config.FEAT_QUESTIONAIRE_CODE, layerCode)
		return c.JSON(errCode, helpers.WebResponseError(errMessage.Error(), code))
	}

	var questionRespose = CoreToResponseList(results)
	code := helpers.GenerateCodeResponse(http.StatusOK, config.FEAT_QUESTIONAIRE_CODE, config.RESPONSE_SUCCESS_CODE)
	return c.JSON(http.StatusOK, helpers.WebResponseSuccess("[success] read data", code, questionRespose))
}
