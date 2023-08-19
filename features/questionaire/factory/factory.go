package factory

import (
	_questionaireData "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/data"
	_questionaireHandler "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"
	_questionaireService "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *_questionaireHandler.QuestionaireHandler {
	repo := _questionaireData.New(db)
	service := _questionaireService.New(repo)
	handler := _questionaireHandler.New(service)

	return handler
}
