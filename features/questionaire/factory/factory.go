package factory

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_questionaireData "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/data"
	_questionaireHandler "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"
	_questionaireService "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/service"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg *config.AppConfig) *_questionaireHandler.QuestionaireHandler {
	repo := _questionaireData.New(db)
	service := _questionaireService.New(repo, cfg)
	handler := _questionaireHandler.New(service)

	return handler
}
