package data

import (
	"time"

	_patientData "github.com/dr-ariawan-s-project/api-drariawan/features/patient/data"
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
	UrlVideo           string
	Section            string
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
	ID            string
	PatientId     string
	CodeAttempt   string
	NotesAttempt  string
	Score         int
	AiAccuracy    float64
	AiProbability float64
	AiDiagnosis   string
	Diagnosis     string
	Feedback      string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Patient       _patientData.Patient
}

func (TestAttempt) TableName() string {
	return "test_attempt"
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
		UrlVideo:    dataModel.UrlVideo,
		Section:     dataModel.Section,
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

func AttempCoreToModel(core questionaire.CoreAttempt) TestAttempt {
	return TestAttempt{
		ID:          core.Id,
		PatientId:   core.PatientId,
		CodeAttempt: core.CodeAttempt,
	}
}

func (data TestAttempt) ModelToCore() questionaire.CoreAttempt {
	return questionaire.CoreAttempt{
		Id:            data.ID,
		PatientId:     data.PatientId,
		CodeAttempt:   data.CodeAttempt,
		NotesAttempt:  data.NotesAttempt,
		Score:         data.Score,
		AIAccuracy:    data.AiAccuracy,
		AIProbability: data.AiProbability,
		AIDiagnosis:   data.AiDiagnosis,
		Diagnosis:     data.Diagnosis,
		Feedback:      data.Feedback,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
		Patient: questionaire.Patient{
			ID:             data.Patient.ID,
			Name:           data.Patient.Name,
			Email:          data.Patient.Email,
			Gender:         data.Patient.Gender,
			MarriageStatus: data.Patient.MarriageStatus,
		},
	}
}

func ListAttemptModelToCore(data []TestAttempt) []questionaire.CoreAttempt {
	var result []questionaire.CoreAttempt
	for _, v := range data {
		result = append(result, v.ModelToCore())
	}
	return result
}

func (data Answer) ModelToCore() questionaire.CoreAnswer {
	return questionaire.CoreAnswer{
		Id:          data.ID,
		AttemptId:   data.AttemptId,
		QuestionId:  data.QuestionId,
		Description: data.Description,
		Score:       data.Score,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Question: questionaire.Core{
			Id:          data.Question.ID,
			Type:        data.Question.Type,
			Question:    data.Question.Question,
			Description: data.Question.Description,
			UrlVideo:    data.Question.UrlVideo,
			Section:     data.Question.Section,
			Goto:        data.Question.Goto,
		},
	}
}

func ListAnswerModelToCore(data []Answer) []questionaire.CoreAnswer {
	var result []questionaire.CoreAnswer
	for _, v := range data {
		result = append(result, v.ModelToCore())
	}
	return result
}
