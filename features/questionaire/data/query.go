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
