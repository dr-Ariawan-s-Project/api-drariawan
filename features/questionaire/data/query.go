package data

import (
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

// CountTestAttemp implements questionaire.QuestionaireDataInterface.
func (repo *questionaireQuery) CountTestAttempt(patientId string) (int, error) {
	var countAttemp int64
	tx := repo.db.Model(&TestAttempt{}).Where("patient_id = ?", patientId).Count(&countAttemp)
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
