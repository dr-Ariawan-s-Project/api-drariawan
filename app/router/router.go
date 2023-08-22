package router

import (
	_questionaireFactory "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/factory"
	_questionaireAPI "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type appsFactory struct {
	questionaireHandler *_questionaireAPI.QuestionaireHandler
}

func InitRouter(db *gorm.DB, e *echo.Echo) {
	sysRoute := appsFactory{
		questionaireHandler: _questionaireFactory.New(db),
	}

	v1 := e.Group("/v1")
	v1Questioner := v1.Group("/questioner")
	v1Questioner.GET("", sysRoute.questionaireHandler.GetAllQuestion)
}
