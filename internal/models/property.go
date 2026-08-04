package models

type Property struct {
	BaseModel
	OwnerID           string `gorm:"type:char(36);not null"`
	Title             string
	Type              string
	Description       string
	Province          string
	District          string
	Ward              string
	Address           string
	Latitude          float64
	Longitude         float64
	Area              float64
	MaxPeople         int
	NumberOfBedrooms  int
	NumberOfBathrooms int
	ExtraRoomInfos    string
	Price             float64
	ElectricityPrice  float64
	WaterPrice        float64
	Status            string          `gorm:"type:enum('available','reserved','rented', 'hidden', 'maintenance')"`
	Images            []PropertyImage `gorm:"foreignKey:PropertyID"`
	Videos            []PropertyVideo `gorm:"foreignKey:PropertyID"`
	Utilities         []Utility       `gorm:"many2many:property_utilities;"`
}
