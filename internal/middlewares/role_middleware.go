package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireRoles is a middleware function that checks if the user's role, extracted from the JWT token, matches any of the allowed roles specified in the roles parameter. If the user's role is not in the allowed roles, it aborts the request and returns a forbidden response.
func RequireRoles(
	roles ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"success": false,
					"message": "Forbidden",
				},
			)
			return
		}

		currentRole := role.(string)

		for _, allowedRole := range roles {
			if currentRole == allowedRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{
				"success": false,
				"message": "Permission denied",
			},
		)
	}
}
