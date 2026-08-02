package middlewares

import (
	"os"
	"strings"

	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a middleware function that checks for the presence of a valid JWT token in the Authorization header of incoming requests. If the token is valid, it extracts the user's role from the token claims and sets it in the request context for further processing. If the token is missing or invalid, it aborts the request and returns an unauthorized response.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.HandleError(c, apperrors.Unauthorized)
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.HandleError(c, apperrors.InvalidToken)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &utils.JwtClaims{}

		accessToken, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrTokenSignatureInvalid
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !accessToken.Valid {
			utils.HandleError(c, apperrors.InvalidToken)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Role)

		c.Next()
	}
}