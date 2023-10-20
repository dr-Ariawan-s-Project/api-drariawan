package service

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"
	"github.com/dr-ariawan-s-project/api-drariawan/features/patient"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
)

type dashboardService struct {
	questionaireServ questionaire.QuestionaireServiceInterface
	patientServ      patient.PatientServiceInterface
}

func New(questionerServ questionaire.QuestionaireServiceInterface, patientServ patient.PatientServiceInterface) dashboard.DashboardServiceInterface {
	return &dashboardService{
		questionaireServ: questionerServ,
		patientServ:      patientServ,
	}
}

// GetDashboardStatistics implements dashboard.DashboardServiceInterface.
func (service *dashboardService) GetDashboardStatistics() (dashboard.DashboardCore, error) {
	var dashboardData dashboard.DashboardCore
	questAttemptCount, errQuestAttempt := service.questionaireServ.CountQuestionerAttempt()

	// get data from status validated
	questAttemptNeedAssess, errQuestAttemptNeedAssess := service.questionaireServ.CountAttemptByStatusAssessment(config.QUESTIONER_ATTEMPT_STATUS_VALIDATED)

	// get data from this month
	t := time.Now()
	questAttemptMonth, errQuestAttemptMonth := service.questionaireServ.CountAttemptByMonth(int(t.Month()))

	//get data all patient
	patientCount, errPatientCount := service.patientServ.CountAllPatient()

	if errQuestAttempt != nil || errQuestAttemptNeedAssess != nil || errQuestAttemptMonth != nil || errPatientCount != nil {
		return dashboardData, errQuestAttempt
	}
	dashboardData.AllQuestioner = questAttemptCount
	dashboardData.NeedAssessQuestioner = questAttemptNeedAssess
	dashboardData.MonthQuestioner = questAttemptMonth
	dashboardData.AllPatient = patientCount

	return dashboardData, nil
}
