package models

type Utility struct {
	BaseModel
	Code       string     `gorm:"type:varchar(100);not null;unique"`
	Properties []Property `gorm:"many2many:property_utilities;"`
}
