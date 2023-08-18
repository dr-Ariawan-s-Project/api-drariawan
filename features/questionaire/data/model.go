package data

import (
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Type               string
	Question           string
	Description        string
	Choices            []Choice   `gorm:"foreignKey:QuestionId"`
	GotoQuestion       []Question `gorm:"foreignKey:goto;references:ID"`
	GotoChoiceQuestion []Choice   `gorm:"foreignKey:goto;references:ID"`
	Goto               *uint
}

type Choice struct {
	gorm.Model
	QuestionId uint
	Option     string
	Slugs      string
	Score      int
	Goto       *uint
}

func ModelChoiceToCore(dataModel Choice) questionaire.Choice {
	return questionaire.Choice{
		Id:         dataModel.ID,
		QuestionId: dataModel.QuestionId,
		Option:     dataModel.Option,
		Slugs:      dataModel.Slugs,
		Score:      dataModel.Score,
		Goto:       dataModel.Goto,
	}
}

func ModelChoiceToCoreList(dataModel []Choice) []questionaire.Choice {
	var choiceList []questionaire.Choice
	for _, v := range dataModel {
		choiceList = append(choiceList, ModelChoiceToCore(v))
	}
	return choiceList
}

func ModelToCore(dataModel Question) questionaire.Core {
	return questionaire.Core{
		Id:          dataModel.ID,
		Type:        dataModel.Type,
		Question:    dataModel.Question,
		Description: dataModel.Description,
		Goto:        dataModel.Goto,
		Choices:     ModelChoiceToCoreList(dataModel.Choices),
	}
}

func ModelToCoreList(dataModel []Question) []questionaire.Core {
	var coreList []questionaire.Core
	for _, v := range dataModel {
		coreList = append(coreList, ModelToCore(v))
	}
	return coreList
}
