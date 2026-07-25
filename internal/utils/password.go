package utils

import "golang.org/x/crypto/bcrypt"

// Hash generates a bcrypt hash of the password.
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(bytes), err
}

// Compare checks if the provided password matches the hashed password.
func Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err == nil
}
