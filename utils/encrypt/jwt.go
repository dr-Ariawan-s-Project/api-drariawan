package encrypt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(t interface{}) (int, string, error) {
	user := t.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		role := claims["role"].(string)
		return int(userId), role, nil
	}
	return 0, "", fmt.Errorf("token invalid")
}
