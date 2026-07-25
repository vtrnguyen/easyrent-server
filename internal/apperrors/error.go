package apperrors

import "net/http"

type AppError struct {
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	EmailAlreadyExists = &AppError{
		Message: "Email already exists",
		Status:  http.StatusConflict,
	}

	PhoneAlreadyExists = &AppError{
		Message: "Phone number already exists",
		Status:  http.StatusConflict,
	}

	InvalidLoginCredentials = &AppError{
		Message: "Invalid email or phone number or password",
		Status:  http.StatusUnauthorized,
	}

	InternalServer = &AppError{
		Message: "Internal server error",
		Status:  http.StatusInternalServerError,
	}

	RecordNotFound = &AppError{
		Message: "Record not found",
		Status:  http.StatusNotFound,
	}
)
