package models

import "time"

type UserInfo struct {
	UserID         string `gorm:"primaryKey"`
	FullName       string
	AvatarURL      string
	Gender         string `gorm:"type:enum('male','female','other')"`
	Birthday       *time.Time
	Address        string
	IdentityNumber string
	Bio            string
	Occupation     string
}
