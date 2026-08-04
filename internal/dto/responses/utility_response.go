package responses

type UtilityResponse struct {
	ID           string `json:"id"`
	Code         string `json:"code"`
	DisplayName  string `json:"display_name"`
	DisplayOrder int    `json:"display_order"`
}
