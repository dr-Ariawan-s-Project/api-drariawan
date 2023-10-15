package middlewares

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ExtractTokenJWT(e echo.Context) (userId any, userRole string, err error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userId"].(type) {
		case float64:
			userId = claims["userId"].(float64)
		case string:
			userId = claims["userId"].(string)
		default:
			return 0, "", fmt.Errorf("token invalid")
		}
		role := claims["role"].(string)
		return userId, role, nil
	}
	return 0, "", fmt.Errorf("token invalid")
}

/*
func untuk convert id user dari ExtractTokenJWT ke int
*/
func ConvertUserID(id any) int {
	idFloat, ok := id.(float64)
	if !ok {
		log.Println("error convert user id")
	}
	return int(idFloat)
}

/*
func untuk convert id user dari ExtractTokenJWT ke string
*/
func ConvertPatientID(id any) string {
	idStr, ok := id.(string)
	if !ok {
		log.Println("error convert patient id")
	}
	return idStr
}
