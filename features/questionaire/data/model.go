package data

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	ID                 uint
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
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

type TestAttempt struct {
	ID           string
	PatientId    string
	CodeAttempt  string
	NotesAttempt string
	Score        int
	Feedback     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Answer struct {
	ID          string
	AttemptId   string
	QuestionId  uint
	Description string
	Score       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	TestAttempt TestAttempt `gorm:"foreignKey:AttemptId;references:ID"`
	Question    Question    `gorm:"foreignKey:QuestionId;references:ID"`
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

// mapping AnswerCore to model
// generate id answer using uuid
func AnswerCoretoModel(attempId string, data []questionaire.CoreAnswer) []Answer {
	var result []Answer
	for _, v := range data {
		id := uuid.New()
		result = append(result, Answer{
			ID:          id.String(),
			AttemptId:   attempId,
			QuestionId:  v.QuestionId,
			Description: v.Description,
			Score:       v.Score,
		})
	}
	return result
}
