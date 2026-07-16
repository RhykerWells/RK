package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateToken() (string, error) {
	token := make([]byte, sessionTokenLength)
	_, err := rand.Read(token)
	if err != nil {
		return "", fmt.Errorf("failed to read from crypto/rand: %w", err)
	}

	return hex.EncodeToString(token), nil
}

func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
