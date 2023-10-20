package dashboard

type DashboardCore struct {
	AllQuestioner        int
	MonthQuestioner      int
	NeedAssessQuestioner int
	AllPatient           int
}

type DashboardServiceInterface interface {
	GetDashboardStatistics() (DashboardCore, error)
}
