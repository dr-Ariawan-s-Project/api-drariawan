package data

import (
	"errors"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/helpers"
	"gorm.io/gorm"
)

type questionaireQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) questionaire.QuestionaireDataInterface {
	return &questionaireQuery{
		db: db,
	}
}

// CountAllQuestion implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) CountAllQuestion() (int, error) {
	var countAttemp int64
	tx := repo.db.Model(&Question{}).Count(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// CountQuestionerAttempt implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) CountQuestionerAttempt() (int, error) {
	var countAttemp int64
	tx := repo.db.Model(&TestAttempt{}).Count(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// CountTestAttemp implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) CountTestAttempt(patientId string) (dataAttempt questionaire.CoreAttempt, count int, err error) {
	// var countAttemp int64
	var attempt TestAttempt
	tx := repo.db.Where("patient_id = ?", patientId).First(&attempt)
	if tx.Error != nil {
		return questionaire.CoreAttempt{}, 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	if attempt.ID == "" {
		return questionaire.CoreAttempt{}, 0, errors.New(config.DB_ERR_RECORD_NOT_FOUND)
	}
	return attempt.ModelToCore(), 1, nil
}

// CheckCountAttemptAnswer implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) CheckCountAttemptAnswer(patientId string) (int, error) {
	var countAttemp int64
	tx := repo.db.Model(&Answer{}).Select("answers.id").Joins("inner join test_attempt on test_attempt.id = answers.attempt_id").Where("test_attempt.patient_id = ?", patientId).Count(&countAttemp)
	if tx.Error != nil {
		return 0, helpers.CheckQueryErrorMessage(tx.Error)
	}
	return int(countAttemp), nil
}

// InsertTestAttemp implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) InsertTestAttempt(data questionaire.CoreAttempt) error {
	var input = AttempCoreToModel(data)
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return helpers.CheckQueryErrorMessage(tx.Error)
	}
	return nil
}

// FindTestAttempt implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) FindTestAttempt(status string, offset int, limit int) (dataAttempt []questionaire.CoreAttempt, err error) {
	var attemptData []TestAttempt
	txSelect := repo.db.Preload("Patient")
	if status != "" {
		txSelect.Where("status = ?", status).Session(&gorm.Session{})
	}

	tx := txSelect.Order("created_at desc").Offset(offset).Limit(limit).Find(&attemptData)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}

	var attemptAll = ListAttemptModelToCore(attemptData)
	return attemptAll, nil
}

// FindAllAnswerByAttempt implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) FindAllAnswerByAttempt(idAttempt string, offset int, limit int) (dataAnswer []questionaire.CoreAnswer, err error) {
	var answerData []Answer
	tx := repo.db.Preload("Question").Where("attempt_id = ?", idAttempt).Order("question_id asc").Offset(offset).Limit(limit).Find(&answerData)

	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}
	if tx.RowsAffected == 0 {
		return nil, helpers.CheckQueryErrorMessage(errors.New(config.DB_ERR_RECORD_NOT_FOUND))
	}

	var answerAll = ListAnswerModelToCore(answerData)
	return answerAll, nil
}

// InsertAnswer implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) InsertAnswer(idAttempt string, data []questionaire.CoreAnswer) error {
	var input = AnswerCoretoModel(idAttempt, data)
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return helpers.CheckQueryErrorMessage(tx.Error)
	}
	return nil
}

// SelectAll implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) SelectAll() ([]questionaire.Core, error) {
	var questionData []Question
	tx := repo.db.Preload("Choices").Find(&questionData)
	if tx.Error != nil {
		return nil, helpers.CheckQueryErrorMessage(tx.Error)
	}

	var questionCoreAll []questionaire.Core = ModelToCoreList(questionData)
	return questionCoreAll, nil
}
