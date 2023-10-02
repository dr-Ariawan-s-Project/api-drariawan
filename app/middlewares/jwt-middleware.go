package middlewares

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ExtractTokenJWT(e echo.Context) (userId int, userRole string, err error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		role := claims["role"].(string)
		return int(userId), role, nil
	}
	return 0, "", fmt.Errorf("token invalid")
}
