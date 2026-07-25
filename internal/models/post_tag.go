package models

type PostTag struct {
	BaseModel
	TagName string `gorm:"type:text;not null"`
	Posts   []Post `gorm:"many2many:post_tag_relations;"`
}
