package handler

import (
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
)

type QuestionResponse struct {
	Id              uint             `json:"id"`
	Type            string           `json:"type"`
	Question        string           `json:"question"`
	Description     string           `json:"description"`
	UrlVideo        string           `json:"url_video"`
	Section         string           `json:"section"`
	ChoiceResponses []ChoiceResponse `json:"choices"`
	Goto            *uint            `json:"goto"`
}

type ChoiceResponse struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"question_id"`
	Option     string `json:"option"`
	Slugs      string `json:"slugs"`
	Score      int    `json:"score"`
	Goto       *uint  `json:"goto"`
}

type QuestionerAttemptResponse struct {
	Id            string          `json:"id"`
	PatientId     string          `json:"patient_id"`
	CodeAttempt   string          `json:"code_attempt"`
	NotesAttempt  string          `json:"notes_attempt"`
	Score         int             `json:"score"`
	AIAccuracy    float64         `json:"ai_accuracy"`
	AIProbability float64         `json:"ai_probability"`
	AIDiagnosis   string          `json:"ai_diagnosis"`
	Diagnosis     string          `json:"diagnosis"`
	Feedback      string          `json:"feedback"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	Patient       PatientResponse `json:"patient"`
}

type PatientResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AnswerResponse struct {
	Id          string           `json:"id"`
	AttemptId   string           `json:"attempt_id"`
	QuestionId  uint             `json:"question_id"`
	Description string           `json:"description"`
	Score       int              `json:"score"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Question    QuestionResponse `json:"question"`
}

func CoreChoicesToResponse(dataCore []questionaire.Choice) []ChoiceResponse {
	var response []ChoiceResponse
	for _, v := range dataCore {
		response = append(response, ChoiceResponse{
			Id:         v.Id,
			QuestionId: v.QuestionId,
			Option:     v.Option,
			Slugs:      v.Slugs,
			Score:      v.Score,
			Goto:       v.Goto,
		})
	}
	return response
}

func CoreToResponse(dataCore questionaire.Core) QuestionResponse {
	return QuestionResponse{
		Id:              dataCore.Id,
		Type:            dataCore.Type,
		Question:        dataCore.Question,
		Description:     dataCore.Description,
		UrlVideo:        dataCore.UrlVideo,
		Section:         dataCore.Section,
		ChoiceResponses: CoreChoicesToResponse(dataCore.Choices),
		Goto:            dataCore.Goto,
	}
}

func CoreToResponseList(dataCore []questionaire.Core) []QuestionResponse {
	var result []QuestionResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}

func AttemptCoreToResponse(dataCore questionaire.CoreAttempt) QuestionerAttemptResponse {
	return QuestionerAttemptResponse{
		Id:            dataCore.Id,
		PatientId:     dataCore.PatientId,
		CodeAttempt:   dataCore.CodeAttempt,
		NotesAttempt:  dataCore.NotesAttempt,
		Score:         dataCore.Score,
		AIAccuracy:    dataCore.AIAccuracy,
		AIProbability: dataCore.AIProbability,
		AIDiagnosis:   dataCore.AIDiagnosis,
		Diagnosis:     dataCore.Diagnosis,
		Feedback:      dataCore.Feedback,
		Status:        dataCore.Status,
		CreatedAt:     dataCore.CreatedAt,
		UpdatedAt:     dataCore.UpdatedAt,
		Patient: PatientResponse{
			ID:    dataCore.Patient.ID,
			Name:  dataCore.Patient.Name,
			Email: dataCore.Patient.Email,
		},
	}
}

func ListAttemptCoreToResponse(dataCore []questionaire.CoreAttempt) []QuestionerAttemptResponse {
	var result []QuestionerAttemptResponse
	for _, v := range dataCore {
		result = append(result, AttemptCoreToResponse(v))
	}
	return result
}

func AnswerCoreToResponse(dataCore questionaire.CoreAnswer) AnswerResponse {
	return AnswerResponse{
		Id:          dataCore.Id,
		AttemptId:   dataCore.AttemptId,
		QuestionId:  dataCore.QuestionId,
		Description: dataCore.Description,
		Score:       dataCore.Score,
		CreatedAt:   dataCore.CreatedAt,
		UpdatedAt:   dataCore.UpdatedAt,
		Question: QuestionResponse{
			Id:       dataCore.Question.Id,
			Type:     dataCore.Question.Type,
			Question: dataCore.Question.Question,
			Section:  dataCore.Question.Section,
		},
	}
}

func ListAnswerCoreToResponse(dataCore []questionaire.CoreAnswer) []AnswerResponse {
	var result []AnswerResponse
	for _, v := range dataCore {
		result = append(result, AnswerCoreToResponse(v))
	}
	return result
}
