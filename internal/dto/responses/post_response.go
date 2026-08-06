package responses

import "time"

type PostResponse struct {
	ID            string     `json:"id"`
	PropertyID    string     `json:"property_id"`
	AuthorID      string     `json:"author_id"`
	Title         string     `json:"title"`
	ContentType   string     `json:"content_type"`
	Content       string     `json:"content"`
	Status        string     `json:"status"`
	PublishedAt   *time.Time `json:"published_at"`
	ExpiresAt     *time.Time `json:"expires_at"`
	PropertyTitle string     `json:"property_title"`
	CreatedAt     time.Time  `json:"created_at"`
}
