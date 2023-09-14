package service

import (
	"errors"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/dr-ariawan-s-project/api-drariawan/features/questionaire"
	"github.com/dr-ariawan-s-project/api-drariawan/utils/encrypt"
)

type questionaireService struct {
	questionaireData questionaire.QuestionaireDataInterface
	cfg              *config.AppConfig
}

func New(repo questionaire.QuestionaireDataInterface, cfg *config.AppConfig) questionaire.QuestionaireServiceInterface {
	return &questionaireService{
		questionaireData: repo,
		cfg:              cfg,
	}
}

// InsertAnswer implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) InsertAnswer(codeAttempt string, data []questionaire.CoreAnswer) error {
	if codeAttempt == "" {
		return errors.New(config.VAL_InvalidValidation)
	}
	// decrypt codeAttempt to idAttempt
	idAttempt := encrypt.DecryptText(codeAttempt, service.cfg.AES_GCM_SECRET)

	err := service.questionaireData.InsertAnswer(idAttempt, data)
	return err
}

// GetAll implements questionaire.QuestionaireServiceInterface.
func (service *questionaireService) GetAll() ([]questionaire.Core, error) {
	return service.questionaireData.SelectAll()
}
