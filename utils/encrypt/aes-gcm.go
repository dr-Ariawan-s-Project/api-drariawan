package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func EncryptText(plaintext string, keySecret string) string {
	key := []byte(keySecret)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("encrypt err : %s\n", err.Error())
		return ""
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("encrypt err : %s\n", err.Error())
		return ""
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Printf("encrypt err : %s\n", err.Error())
		return ""
	}

	plaintextBytes := []byte(plaintext)
	ciphertext := aesgcm.Seal(nil, nonce, plaintextBytes, nil)

	// add nonce to ciphertext
	ciphertextWithNonce := append(nonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(ciphertextWithNonce)
}

func DecryptText(ciphertextWithNonce string, keySecret string) string {
	key := []byte(keySecret)

	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextWithNonce)
	if err != nil {
		log.Printf("decrypt err : %s\n", err.Error())
		return ""
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := 12 // fill with your nonce size (usually 12 byte)
	nonce, ciphertext := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("decrypt err : %s\n", err.Error())
		return ""
	}

	// Create a new GCM (Galois Counter Mode) cipher
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("decrypt err : %s\n", err.Error())
		return ""
	}

	// Decrypt the ciphertext using the GCM cipher
	plaintextBytes, err := aesgcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		log.Printf("decrypt err : %s\n", err.Error())
		return ""
	}

	// Convert the decrypted plaintext byte slice to a string
	plaintext := string(plaintextBytes)

	// Return the decrypted plaintext string
	return plaintext
}
