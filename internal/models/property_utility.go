package models

type PropertyUtility struct {
	PropertyID string   `gorm:"type:char(36);primaryKey"`
	UtilityID  string   `gorm:"type:char(36);primaryKey"`
	Property   Property `gorm:"foreignKey:PropertyID"`
	Utility    Utility  `gorm:"foreignKey:UtilityID"`
}
