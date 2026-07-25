package requests

type RegisterRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}
