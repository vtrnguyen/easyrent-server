package requests

type PostRequest struct {
	PropertyID  string `json:"property_id" binding:"required"`
	Title       string `json:"title" binding:"required,max=255"`
	ContentType string `json:"content_type" binding:"required,oneof=plain_text markdown"`
	Content     string `json:"content" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=draft pending_review published hidden expired"`
}
