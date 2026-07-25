package models

type PropertyType struct {
	BaseModel
	Code       string     `gorm:"type:enum('house','rental_room','apartment','flat');not null;unique"`
	Properties []Property `gorm:"foreignKey:PropertyTypeID"`
}
