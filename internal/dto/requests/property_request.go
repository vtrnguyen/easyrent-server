package requests

type PropertyRequest struct {
	OwnerID           string   `form:"owner_id"`
	Title             string   `form:"title" binding:"required"`
	Type              string   `form:"type" binding:"required"`
	Description       string   `form:"description" binding:"required"`
	Province          string   `form:"province" binding:"required"`
	District          string   `form:"district" binding:"required"`
	Ward              string   `form:"ward" binding:"required"`
	Address           string   `form:"address" binding:"required"`
	Latitude          float64  `form:"latitude" binding:"required"`
	Longitude         float64  `form:"longitude" binding:"required"`
	Area              float64  `form:"area" binding:"required"`
	MaxPeople         int      `form:"max_people" binding:"required"`
	NumberOfBedrooms  int      `form:"number_of_bedrooms" binding:"required"`
	NumberOfBathrooms int      `form:"number_of_bathrooms" binding:"required"`
	ExtraRoomInfos    string   `form:"extra_room_infos"`
	Price             float64  `form:"price" binding:"required"`
	ElectricityPrice  float64  `form:"electricity_price" binding:"required"`
	WaterPrice        float64  `form:"water_price" binding:"required"`
	Status            string   `form:"status" binding:"required,oneof=available reserved rented hidden maintenance"`
	UtilityIDs        []string `form:"utilities"`
}
