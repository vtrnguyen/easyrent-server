package requests

import "mime/multipart"

type CreateUserRequest struct {
	Email          string                `form:"email" binding:"required,email"`
	PhoneNumber    string                `form:"phone_number"`
	Role           string                `form:"role"`
	Status         string                `form:"status"`
	FullName       string                `form:"full_name"`
	Gender         string                `form:"gender"`
	Birthday       string                `form:"birthday"`
	Occupation     string                `form:"occupation"`
	IdentityNumber string                `form:"identity_number"`
	Address        string                `form:"address"`
	Bio            string                `form:"bio"`
	Avatar         *multipart.FileHeader `form:"avatar"`
}

type UpdateUserRequest struct {
	Email          string                `json:"email" binding:"required,email"`
	PhoneNumber    string                `json:"phone_number" binding:"required"`
	Role           string                `json:"role" binding:"required,oneof=tenant landlord admin"`
	Status         string                `json:"status" binding:"omitempty,oneof=active inactive"`
	FullName       string                `json:"fullname" binding:"required"`
	Gender         string                `json:"gender" binding:"required,oneof=male female other"`
	Birthday       string                `json:"birthday" binding:"omitempty,datetime=1900-01-01"`
	Occupation     string                `json:"occupation" binding:"omitempty"`
	IdentityNumber string                `json:"identity_number" binding:"omitempty"`
	Address        string                `json:"address" binding:"omitempty"`
	Bio            string                `json:"bio" binding:"omitempty"`
	Avatar         *multipart.FileHeader `form:"avatar"`
}

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
