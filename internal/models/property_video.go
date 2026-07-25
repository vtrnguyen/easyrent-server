package models

type PropertyVideo struct {
	BaseModel
	PropertyID string   `gorm:"type:char(36);not null"`
	VideoURL   string   `gorm:"type:text;not null"`
	Property   Property `gorm:"foreignKey:PropertyID"`
}
