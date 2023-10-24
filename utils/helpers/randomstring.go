package helpers

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func UUIDGenerate() string {
	id := uuid.New()
	return id.String()
}

func RandomStringAlphabetNumeric() string {
	// Define the character set that includes both alphabet and numeric characters
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random 12-character string
	result := make([]byte, 12)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(result)
}
