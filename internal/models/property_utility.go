package models

type PropertyUtility struct {
	PropertyID string   `gorm:"primaryKey"`
	UtilityID  string   `gorm:"primaryKey"`
	Property   Property `gorm:"foreignKey:PropertyID"`
	Utility    Utility  `gorm:"foreignKey:UtilityID"`
}
