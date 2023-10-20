package service

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"
)

type dashboardService struct {
	dashboardData dashboard.DashboardDataInterface
}

func New(dashboardData dashboard.DashboardDataInterface) dashboard.DashboardServiceInterface {
	return &dashboardService{
		dashboardData: dashboardData,
	}
}

// GetDashboardStatistics implements dashboard.DashboardServiceInterface.
func (service *dashboardService) GetDashboardStatistics() (dashboard.DashboardCore, error) {
	var dashboardResult dashboard.DashboardCore
	questAttemptCount, errQuestAttempt := service.dashboardData.CountQuestionerAttempt()

	// get data from status validated
	questAttemptNeedAssess, errQuestAttemptNeedAssess := service.dashboardData.CountAttemptByStatusAssessment(config.QUESTIONER_ATTEMPT_STATUS_VALIDATED)

	// get data from this month
	t := time.Now()
	questAttemptMonth, errQuestAttemptMonth := service.dashboardData.CountAttemptByMonth(int(t.Month()))

	//get data all patient
	patientCount, errPatientCount := service.dashboardData.CountAllPatient()

	if errQuestAttempt != nil || errQuestAttemptNeedAssess != nil || errQuestAttemptMonth != nil || errPatientCount != nil {
		return dashboardResult, errQuestAttempt
	}
	dashboardResult.AllQuestioner = questAttemptCount
	dashboardResult.NeedAssessQuestioner = questAttemptNeedAssess
	dashboardResult.MonthQuestioner = questAttemptMonth
	dashboardResult.AllPatient = patientCount

	return dashboardResult, nil
}

// GetQuestionerAttemptPerMonth implements dashboard.DashboardServiceInterface.
func (service *dashboardService) GetQuestionerAttemptPerMonth() ([]dashboard.DashboardAttemptCore, error) {
	result, err := service.dashboardData.CountQuestionerAttemptPerMonth()
	return result, err
}
