package models

import "time"

type Post struct {
	BaseModel
	PropertyID      string `gorm:"type:char(36);not null"`
	AuthorID        string `gorm:"type:char(36);not null"`
	Title           string `gorm:"type:varchar(255);not null"`
	MarkdownContent string `gorm:"type:longtext"`
	PlainContent    string `gorm:"type:longtext"`
	Status          string `gorm:"type:enum('draft','pending_review','published','hidden','expired','deleted');default:'draft'"`
	PublishedAt     *time.Time
	ExpiresAt       *time.Time
	Property        Property      `gorm:"foreignKey:PropertyID"`
	Author          User          `gorm:"foreignKey:AuthorID"`
	Tags            []PostTag     `gorm:"many2many:post_tag_relations;"`
	Likes           []PostLike    `gorm:"foreignKey:PostID"`
	Comments        []PostComment `gorm:"foreignKey:PostID"`
}
