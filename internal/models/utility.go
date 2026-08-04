package models

type Utility struct {
	BaseModel
	Code         string `gorm:"type:varchar(100);not null;unique"`
	DisplayName  string `gorm:"type:varchar(255);not null"`
	DisplayOrder int    `gorm:"type:int;not null"`
}
