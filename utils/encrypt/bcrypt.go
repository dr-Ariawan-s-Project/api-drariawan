package encrypt

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil //true means login success
}

func GeneratePassword(password string) string {
	hashed := ""
	if password != "" {
		hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("==== BCRYPT ERROR ==== ", err.Error())
		}
		hashed = string(hashedByte)
	}
	return hashed
}

func ComparePassword(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		log.Println("login compare", err.Error())
		return errors.New("password not match")
	}
	return nil
}
