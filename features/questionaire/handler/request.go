package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"

type AnswerRequest struct {
	CodeAttempt string                `json:"code_attempt"`
	Answer      []AnswerDetailRequest `json:"answer"`
}

type AnswerDetailRequest struct {
	QuestionId  uint   `json:"question_id"`
	Description string `json:"description"`
	Score       int    `json:"score"`
}

type ValidateRequest struct {
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PartnerEmail string `json:"partner_email"`
	As           string `json:"as"`
}

type AssesmentRequest struct {
	Diagnosis string `json:"diagnosis"`
	Feedback  string `json:"feedback"`
	Status    string `json:"status"`
}

func (data AnswerRequest) RequestToCore() []questionaire.CoreAnswer {
	var result []questionaire.CoreAnswer
	for _, v := range data.Answer {
		result = append(result, questionaire.CoreAnswer{
			QuestionId:  v.QuestionId,
			Description: v.Description,
			Score:       v.Score,
		})
	}
	return result
}

func (data AssesmentRequest) RequestToCore() questionaire.CoreAttempt {
	return questionaire.CoreAttempt{
		Diagnosis: data.Diagnosis,
		Feedback:  data.Feedback,
		Status:    data.Status,
	}
}
