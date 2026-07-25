package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

// GenerateAccessToken generates a JWT access token for the given user ID and role.
func GenerateAccessToken(
	userID string,
	role string,
) (string, error) {
	claims := JwtClaims{
		UserId: userID,
		Role:   role,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return accessToken.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}
