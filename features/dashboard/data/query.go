package data

import (
	"github.com/dr-ariawan-s-project/api-drariawan/features/dashboard"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"gorm.io/gorm"
)

type dashboardQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) dashboard.DashboardDataInterface {
	return &dashboardQuery{
		db: db,
	}
}

// CountAttemptByMonth implements questionaire.QuestionaireDataInterface.
func (repo *dashboardQuery) CountAttemptByMonth(month int) (int, error) {
	var countAttemp int64
	tx := repo.db.Raw("select count(id) as count_attempt from test_attempt where MONTH(created_at) = ? and deleted_at is null", month).Scan(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// CountAttemptByStatusAssessment implements questionaire.QuestionaireDataInterface.
func (repo *dashboardQuery) CountAttemptByStatusAssessment(status string) (int, error) {
	var countAttemp int64
	tx := repo.db.Raw("select count(id) as count_attempt from test_attempt where status != ? OR diagnosis IS NULL and deleted_at is null", status).Scan(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// CountQuestionerAttempt implements questionaire.QuestionaireDataInterface.
func (repo *dashboardQuery) CountQuestionerAttempt() (int, error) {
	var countAttemp int64
	tx := repo.db.Raw("select count(id) as count_attempt from test_attempt where deleted_at is null").Scan(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// for dashboard
// CountAllPatient implements patient.PatientDataInterface.
func (repo *dashboardQuery) CountAllPatient() (int, error) {
	var countAttemp int64
	tx := repo.db.Raw("select count(id) as count_attempt from patients where deleted_at is null").Scan(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// CountQuestionerAttemptPerMonth implements dashboard.DashboardDataInterface.
func (repo *dashboardQuery) CountQuestionerAttemptPerMonth() ([]dashboard.DashboardAttemptCore, error) {
	query := `WITH months_list AS
	(
		select 1 as Number UNION
		select 2 UNION
		select 3 UNION
		select 4 UNION
		select 5 UNION
		select 6 UNION
		select 7 UNION
		select 8 UNION
		select 9 UNION
		select 10 UNION
		select 11 UNION
		select 12 
	), count_attempt_month as (
	  SELECT month(ta.created_at) as month_created,
			 count(ta.id) AS attempt_count
	  from test_attempt ta
	  WHERE ta.created_at BETWEEN STR_TO_DATE(CONCAT(YEAR(CURDATE()), '-01-01'), '%Y-%m-%d') and DATE_ADD(STR_TO_DATE(CONCAT(YEAR(CURDATE()), '-01-01'), '%Y-%m-%d'), INTERVAL 1 YEAR) and ta.deleted_at is NULL
	  group by month(created_at)
	)
	select months_list.Number as month, ifnull(cm.attempt_count,0) as count_attempt
	 from months_list
	 left join count_attempt_month cm on months_list.Number = cm.month_created`
	var dataAttemp []DashboardQuestioner
	tx := repo.db.Raw(query).Scan(&dataAttemp)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}
	result := ListDashboardAttemptModelToCore(dataAttemp)
	return result, nil
}
