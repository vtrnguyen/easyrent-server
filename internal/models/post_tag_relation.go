package models

type PostTagRelation struct {
	PostID string `gorm:"primaryKey"`
	TagID  string `gorm:"primaryKey"`
}
