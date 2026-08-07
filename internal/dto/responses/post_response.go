package responses

import "time"

type PostResponse struct {
	ID              string     `json:"id"`
	PropertyID      string     `json:"property_id"`
	AuthorID        string     `json:"author_id"`
	Title           string     `json:"title"`
	ContentType     string     `json:"content_type"`
	Content         string     `json:"content"`
	Status          string     `json:"status"`
	PublishedAt     *time.Time `json:"published_at"`
	ExpiresAt       *time.Time `json:"expires_at"`
	PropertyTitle   string     `json:"property_title"`
	PropertyPrice   float64    `json:"property_price"`
	PropertyArea    float64    `json:"property_area"`
	PropertyAddress string     `json:"property_address"`
	ThumbnailURL    string     `json:"thumbnail_url"`
	CreatedAt       time.Time  `json:"created_at"`
}

type PostCommentResponse struct {
	ID              string                `json:"id"`
	UserID          string                `json:"user_id"`
	UserName        string                `json:"user_name"`
	UserAvatarURL   string                `json:"user_avatar_url"`
	ParentCommentID *string               `json:"parent_comment_id"`
	Content         string                `json:"content"`
	CreatedAt       time.Time             `json:"created_at"`
	Children        []PostCommentResponse `json:"children"`
}

type PostSocialResponse struct {
	Liked     bool  `json:"liked"`
	LikeCount int64 `json:"like_count"`
}
