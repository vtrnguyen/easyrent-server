package handlers

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new instance of AuthHandler with the necessary dependencies.
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

// Register handles the user registration process. It validates the request, calls the AuthService to register the user, and returns an appropriate response.
func (h *AuthHandler) Register(c *gin.Context) {
	var req requests.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	data, err := h.authService.Register(req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Registered successfully", data)
}

// Login handles the user login process. It validates the request, calls the AuthService to authenticate the user, and returns an appropriate response.
func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	data, err := h.authService.Login(req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Logged in successfully", data)
}
