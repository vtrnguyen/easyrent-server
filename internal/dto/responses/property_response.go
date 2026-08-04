package responses

import "time"

type PropertyImageResponse struct {
	ID           string `json:"id"`
	ImageURL     string `json:"image_url"`
	IsThumbnail  bool   `json:"is_thumbnail"`
	DisplayOrder int    `json:"display_order"`
}

type PropertyVideoResponse struct {
	ID       string `json:"id"`
	VideoURL string `json:"video_url"`
}

type PropertyResponse struct {
	ID                string                  `json:"id"`
	OwnerID           string                  `json:"owner_id"`
	Title             string                  `json:"title"`
	Type              string                  `json:"type"`
	Description       string                  `json:"description"`
	Province          string                  `json:"province"`
	District          string                  `json:"district"`
	Ward              string                  `json:"ward"`
	Address           string                  `json:"address"`
	Latitude          float64                 `json:"latitude"`
	Longitude         float64                 `json:"longitude"`
	Area              float64                 `json:"area"`
	MaxPeople         int                     `json:"max_people"`
	NumberOfBedrooms  int                     `json:"number_of_bedrooms"`
	NumberOfBathrooms int                     `json:"number_of_bathrooms"`
	ExtraRoomInfos    string                  `json:"extra_room_infos"`
	Price             float64                 `json:"price"`
	ElectricityPrice  float64                 `json:"electricity_price"`
	WaterPrice        float64                 `json:"water_price"`
	Status            string                  `json:"status"`
	Images            []PropertyImageResponse `json:"images"`
	Videos            []PropertyVideoResponse `json:"videos"`
	Utilities         []string                `json:"utilities"`
	CreatedAt         time.Time               `json:"created_at"`
	UpdatedAt         time.Time               `json:"updated_at"`
}
