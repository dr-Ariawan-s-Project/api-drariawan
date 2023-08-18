package service

import "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"

type questionaireService struct {
	questionaireData questionaire.QuestionaireDataInterface
}

func New(repo questionaire.QuestionaireDataInterface) questionaire.QuestionaireServiceInterface {
	return &questionaireService{
		questionaireData: repo,
	}
}

// GetAll implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetAll() ([]questionaire.Core, error) {
	return service.questionaireData.SelectAll()
}
