package encrypt

import (
	"encoding/base64"
	"fmt"
)

func Encrypt(plaintext []byte) (string, error) {
	if plaintext == nil {
		return "", fmt.Errorf("Error encrypt text empty")
	}
	return base64.URLEncoding.EncodeToString(plaintext), nil
}

func Decrypt(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", fmt.Errorf("Error dencrypt text empty")
	}
	decodeValue, err := base64.URLEncoding.DecodeString(ciphertext)
	return string(decodeValue), err
}
