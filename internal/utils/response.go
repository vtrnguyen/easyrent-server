package utils

import (
	"easyrent-server/internal/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    string      `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Success sends a successful JSON response with the provided message and data.
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error sends an error JSON response with the provided status, message, and errors.
func Error(
	c *gin.Context,
	status int,
	message string,
	errors interface{},
) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// HandleError handles errors and sends an appropriate JSON response based on the error type.
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if appErr, ok := err.(*apperrors.AppError); ok {
		Error(
			c,
			appErr.Status,
			appErr.Message,
			nil,
		)
		return
	}

	Error(
		c,
		http.StatusInternalServerError,
		apperrors.InternalServer.Message,
		nil,
	)
}

// HandleValidationError handles validation errors and sends a JSON response with the validation error details.
func HandleValidationError(c *gin.Context, err error) {
	Error(
		c,
		http.StatusBadRequest,
		"Validation failed",
		ParseValidationErrors(err),
	)
}
