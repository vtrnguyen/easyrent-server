package models

import "time"

type PostLike struct {
	UserID    string `gorm:"primaryKey"`
	PostID    string `gorm:"primaryKey"`
	CreatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
	Post      Post `gorm:"foreignKey:PostID"`
}
