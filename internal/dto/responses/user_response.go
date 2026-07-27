package responses

import "time"

type UserResponse struct {
	ID             string     `json:"id"`
	Email          string     `json:"email"`
	PhoneNumber    string     `json:"phone_number"`
	Role           string     `json:"role"`
	Status         string     `json:"status"`
	EmailVerified  bool       `json:"email_verified"`
	FullName       string     `json:"full_name"`
	AvatarURL      string     `json:"avatar_url"`
	Gender         string     `json:"gender"`
	Birthday       *time.Time `json:"birthday"`
	Address        string     `json:"address"`
	IdentityNumber string     `json:"identity_number"`
	Bio            string     `json:"bio"`
	Occupation     string     `json:"occupation"`
	LastLoginAt    *time.Time `json:"last_login_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
