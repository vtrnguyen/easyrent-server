package requests

type UpdateMeRequest struct {
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	FullName       string `json:"fullname"`
	AvatarURL      string `json:"avatar_url"`
	Gender         string `json:"gender"`
	Birthday       string `json:"birthday"`
	Address        string `json:"address"`
	IdentityNumber string `json:"identity_number"`
	Bio            string `json:"bio"`
	Occupation     string `json:"occupation"`
}
