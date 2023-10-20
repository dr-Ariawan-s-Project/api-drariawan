package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"

type DashboardQuestionerResponse struct {
	AllQuestioner   int `json:"questioner_all"`
	NeedAssess      int `json:"questioner_need_assess"`
	MonthQuestioner int `json:"questioner_this_month"`
	AllPatient      int `json:"patient_all"`
}

type DashboardAttemptResponse struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

func ListDashboardAttemptCoreToResponse(data []dashboard.DashboardAttemptCore) []DashboardAttemptResponse {
	var result []DashboardAttemptResponse
	for _, v := range data {
		result = append(result, DashboardAttemptResponse{
			Month: v.Month,
			Count: v.Count,
		})
	}
	return result
}
