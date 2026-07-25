package auth

import "time"

type AuthResponse struct {
	AccessToken string     `json:"access_token"`
	UserID      string     `json:"user_id"`
	Role        string     `json:"role"`
	Status      string     `json:"status"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}
