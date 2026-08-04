package models

type PropertyImage struct {
	BaseModel
	PropertyID   string `gorm:"type:char(36);not null"`
	ImageURL     string
	IsThumbnail  bool
	DisplayOrder int
	Property     Property `gorm:"foreignKey:PropertyID"`
}
