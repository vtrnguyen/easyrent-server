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

	InvalidPassword = &AppError{
		Message: "Invalid password",
		Status:  http.StatusBadRequest,
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

	Unauthorized = &AppError{
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
	}

	InvalidToken = &AppError{
		Message: "Invalid token",
		Status:  http.StatusUnauthorized,
	}

	Forbidden = &AppError{
		Message: "Forbidden",
		Status:  http.StatusForbidden,
	}

	PasswordConfirmationNotMatch = &AppError{
		Message: "Password confirmation does not match",
		Status:  http.StatusBadRequest,
	}
)
