package responses

import "time"

type AuthResponse struct {
	AccessToken string     `json:"access_token"`
	UserID      string     `json:"user_id"`
	FullName    string     `json:"full_name"`
	AvatarURL   string     `json:"avatar_url"`
	Role        string     `json:"role"`
	Status      string     `json:"status"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}
