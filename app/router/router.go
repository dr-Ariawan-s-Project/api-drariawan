package router

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_authFactory "github.com/dr-ariawan-s-project/api-drariawan/features/auth/factory"
	_authAPI "github.com/dr-ariawan-s-project/api-drariawan/features/auth/handler"
	_patientFactory "github.com/dr-ariawan-s-project/api-drariawan/features/patient/factory"
	_patientAPI "github.com/dr-ariawan-s-project/api-drariawan/features/patient/handler"
	_questionaireFactory "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/factory"
	_questionaireAPI "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type appsFactory struct {
	questionaireHandler *_questionaireAPI.QuestionaireHandler
	authHandler         *_authAPI.AuthHandler
	patientHandler      *_patientAPI.PatientHandler
}

func InitRouter(db *gorm.DB, e *echo.Echo, cfg *config.AppConfig) {
	sysRoute := appsFactory{
		questionaireHandler: _questionaireFactory.New(db, cfg),
		authHandler:         _authFactory.New(db, cfg),
		patientHandler:      _patientFactory.New(db, cfg),
	}
	e.POST("/login", sysRoute.authHandler.Login)

	v1 := e.Group("/v1")
	v1Questioner := v1.Group("/questioner")
	v1Questioner.GET("", sysRoute.questionaireHandler.GetAllQuestion)
	v1Questioner.POST("", sysRoute.questionaireHandler.AddAnswer)

	v1Patient := v1.Group("/patients")
	v1Patient.POST("", sysRoute.patientHandler.AddPatient)
	v1Patient.GET("", sysRoute.patientHandler.GetAll)
	v1Patient.GET("/:patient_id", sysRoute.patientHandler.GetById)
	v1Patient.PUT("/:patient_id", sysRoute.patientHandler.EditPatient)
	v1Patient.DELETE("/:patient_id", sysRoute.patientHandler.DeleteById)
}
