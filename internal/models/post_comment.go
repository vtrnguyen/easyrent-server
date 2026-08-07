package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostComment struct {
	ID              string `gorm:"type:char(36);primaryKey"`
	UserID          string `gorm:"type:char(36);not null"`
	PostID          string `gorm:"type:char(36);not null"`
	ParentCommentID *string
	Content         string `gorm:"type:text;not null"`
	CreatedAt       time.Time
	User            User         `gorm:"foreignKey:UserID"`
	Post            Post         `gorm:"foreignKey:PostID"`
	ParentComment   *PostComment `gorm:"foreignKey:ParentCommentID"`
}

func (comment *PostComment) BeforeCreate(tx *gorm.DB) error {
	if comment.ID == "" {
		comment.ID = uuid.New().String()
	}
	return nil
}
