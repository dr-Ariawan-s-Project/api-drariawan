package handler

import "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"

type QuestionResponse struct {
	Id              uint             `json:"id"`
	Type            string           `json:"type"`
	Question        string           `json:"question"`
	Description     string           `json:"description"`
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