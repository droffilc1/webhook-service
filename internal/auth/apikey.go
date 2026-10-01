// Package auth implements API key generation
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func Generate() (key, hash string) {
	key = "whk_" + rand.Text()
	hash = Hash(key)

	return key, hash
}

func Hash(key string) string {
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:])
}
