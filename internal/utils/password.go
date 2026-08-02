package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

// Hash generates a bcrypt hash of the password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(bytes), err
}

// ComparePassword checks if the provided password matches the hashed password.
func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err == nil
}

// Generate creates a random password of the specified length.
func GeneratePassword(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte,length)

	for i:=0; i < length; i++ {
		result[i] = characters[
			rand.Intn(len(characters)),
		]

	}

	return string(result)
}
