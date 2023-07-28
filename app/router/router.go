package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello world",
		})
	})
}
