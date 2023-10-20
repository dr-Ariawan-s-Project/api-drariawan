package dashboard

type DashboardCore struct {
	AllQuestioner        int
	MonthQuestioner      int
	NeedAssessQuestioner int
	AllPatient           int
}

type DashboardAttemptCore struct {
	Month string
	Count int
}

type DashboardDataInterface interface {
	CountQuestionerAttemptPerMonth() ([]DashboardAttemptCore, error)
	CountQuestionerAttempt() (int, error)
	CountAttemptByMonth(month int) (int, error)
	CountAttemptByStatusAssessment(status string) (int, error)
	CountAllPatient() (int, error)
}

type DashboardServiceInterface interface {
	GetDashboardStatistics() (DashboardCore, error)
	GetQuestionerAttemptPerMonth() ([]DashboardAttemptCore, error)
}
