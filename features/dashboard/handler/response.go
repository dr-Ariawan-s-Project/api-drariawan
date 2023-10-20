package handler

type DashboardQuestionerResponse struct {
	AllQuestioner   int `json:"questioner_all"`
	NeedAssess      int `json:"questioner_need_assess"`
	MonthQuestioner int `json:"questioner_this_month"`
	AllPatient      int `json:"patient_all"`
}
