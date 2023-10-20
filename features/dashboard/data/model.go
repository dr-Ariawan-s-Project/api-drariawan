package data

import "github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"

type DashboardQuestioner struct {
	Month        int
	CountAttempt int
}

var monthName = map[int]string{
	1:  "januari",
	2:  "februari",
	3:  "maret",
	4:  "april",
	5:  "mei",
	6:  "juni",
	7:  "juli",
	8:  "agustus",
	9:  "september",
	10: "oktober",
	11: "november",
	12: "desember",
}

func (data DashboardQuestioner) ModelToCore() dashboard.DashboardAttemptCore {
	return dashboard.DashboardAttemptCore{
		Month: monthName[data.Month],
		Count: data.CountAttempt,
	}
}

func ListDashboardAttemptModelToCore(data []DashboardQuestioner) []dashboard.DashboardAttemptCore {
	var result []dashboard.DashboardAttemptCore
	for _, v := range data {
		result = append(result, v.ModelToCore())
	}
	return result
}
