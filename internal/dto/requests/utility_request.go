package requests

type UtilityRequest struct {
	Code         string `json:"code" binding:"required"`
	DisplayName  string `json:"display_name" binding:"required"`
	DisplayOrder int    `json:"display_order" binding:"required"`
}
