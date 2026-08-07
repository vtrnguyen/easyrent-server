package requests

type PostCommentRequest struct {
	Content         string  `json:"content" binding:"required,max=2000"`
	ParentCommentID *string `json:"parent_comment_id" binding:"omitempty,uuid"`
}
type RentalRequestRequest struct {
	PropertyID string `json:"property_id" binding:"required,uuid"`
	Message    string `json:"message" binding:"required,max=2000"`
}
