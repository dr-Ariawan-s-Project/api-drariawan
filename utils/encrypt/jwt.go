package encrypt

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dr-ariawan-s-project/api-drariawan/app/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (string, interface{}) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 3 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	useToken, err := token.SignedString([]byte(config.SECRET_JWT))
	if err != nil {
		log.Println(err.Error())
	}
	// log.Println(useToken, "/n", token)
	return useToken, token
}

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

func ExpiredToken() string {
	t := time.Now().Add(time.Hour * 12)
	mont := int(t.Month())
	y := strconv.Itoa(t.Year())
	m := strconv.Itoa(mont)
	d := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())

	if len(m) == 1 {
		m = "0" + m
	}
	if len(d) == 1 {
		d = "0" + d
	}
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(min) == 1 {
		min = "0" + min
	}

	expired := fmt.Sprintf("%s-%s-%s %s:%s", y, m, d, hour, min)
	return expired
}
