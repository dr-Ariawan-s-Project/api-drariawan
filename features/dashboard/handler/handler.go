package handler

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	dashboardService dashboard.DashboardServiceInterface
}

func New(service dashboard.DashboardServiceInterface) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: service,
	}
}

func (handler *DashboardHandler) GetDashboardQuestioner(c echo.Context) error {
	results, err := handler.dashboardService.GetDashboardStatistics()

	if err != nil {
		jsonResponse, httpCode := helpers.WebResponseError(err, config.FEAT_DASHBOARD_CODE)
		return c.JSON(httpCode, jsonResponse)
	}

	var questionerDashboardRespose = DashboardQuestionerResponse{
		AllQuestioner:   results.AllQuestioner,
		NeedAssess:      results.NeedAssessQuestioner,
		MonthQuestioner: results.MonthQuestioner,
		AllPatient:      results.AllPatient,
	}
	mapResponse, httpCode := helpers.WebResponseSuccess("[success] read data", config.FEAT_DASHBOARD_CODE, questionerDashboardRespose)
	return c.JSON(httpCode, mapResponse)
}
