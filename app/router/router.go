package router

import (
	_questionaireData "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/data"
	_questionaireHandler "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"
	_questionaireService "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {

	questionaireData := _questionaireData.New(db)
	questionaireService := _questionaireService.New(questionaireData)
	questionaireHandlerAPI := _questionaireHandler.New(questionaireService)

	e.GET("/questioner", questionaireHandlerAPI.GetAllQuestion)
}
