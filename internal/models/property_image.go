package models

type PropertyImage struct {
	BaseModel
	PropertyID   string
	ImageURL     string
	IsThumbnail  bool
	DisplayOrder int
	Property     Property `gorm:"foreignKey:PropertyID"`
}
