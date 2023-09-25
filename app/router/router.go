package router

import (
	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	_authFactory "github.com/dr-ariawan-s-project/api-drariawan/features/auth/factory"
	_authAPI "github.com/dr-ariawan-s-project/api-drariawan/features/auth/handler"
	_questionaireFactory "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/factory"
	_questionaireAPI "github.com/dr-ariawan-s-project/api-drariawan/features/questionaire/handler"
	_usersFactory "github.com/dr-ariawan-s-project/api-drariawan/features/users/factory"
	_usersAPI "github.com/dr-ariawan-s-project/api-drariawan/features/users/handler"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type appsFactory struct {
	questionaireHandler *_questionaireAPI.QuestionaireHandler
	authHandler         *_authAPI.AuthHandler
	userHandler         *_usersAPI.UserHandler
}

func InitRouter(db *gorm.DB, e *echo.Echo, cfg *config.AppConfig) {
	sysRoute := appsFactory{
		questionaireHandler: _questionaireFactory.New(db, cfg),
		authHandler:         _authFactory.New(db, cfg),
		userHandler:         _usersFactory.New(db),
	}
	e.POST("/login", sysRoute.authHandler.Login)

	v1 := e.Group("/v1")
	v1Questioner := v1.Group("/questioner")
	v1Questioner.GET("", sysRoute.questionaireHandler.GetAllQuestion)
	v1Questioner.POST("", sysRoute.questionaireHandler.AddAnswer)

	// users
	v1User := v1.Group("/user")
	v1User.POST("", sysRoute.userHandler.Insert())
	v1User.PUT("", sysRoute.userHandler.Update(), echojwt.WithConfig(echojwt.Config{SigningMethod: "HS256", SigningKey: []byte(config.InitConfig().JWT_SECRET)}))
	v1User.POST("/deactive", sysRoute.userHandler.Delete())
	v1User.GET("", sysRoute.userHandler.FindById())
	v1User.GET("/list", sysRoute.userHandler.FindAll())
}
