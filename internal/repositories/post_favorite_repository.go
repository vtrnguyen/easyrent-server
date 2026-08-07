package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/models"
)

type PostFavoriteRepository struct{}

// NewPostFavoriteRepository creates a new PostFavoriteRepository instance.
func (r *PostFavoriteRepository) Add(userID, postID string) error {
	favorite := models.PostFavorite{UserID: userID, PostID: postID}

	return database.DB.FirstOrCreate(&favorite, "user_id = ? AND post_id = ?", userID, postID).Error
}

// Remove removes a post from the user's favorites.
func (r *PostFavoriteRepository) Remove(userID, postID string) error {
	return database.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.PostFavorite{}).Error
}

// IDs retrieves the IDs of the user's favorite posts.
func (r *PostFavoriteRepository) IDs(userID string) ([]string, error) {
	ids := make([]string, 0)
	err := database.DB.Model(&models.PostFavorite{}).Where("user_id = ?", userID).Pluck("post_id", &ids).Error

	return ids, err
}

// Search retrieves the user's favorite posts with pagination.
func (r *PostFavoriteRepository) Search(userID string, page, limit int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	query := database.DB.Model(&models.Post{}).
		Joins("JOIN post_favorites ON post_favorites.post_id = posts.id").
		Where("post_favorites.user_id = ? AND posts.status = ?", userID, "published")
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Property").Order("post_favorites.created_at DESC").
		Limit(limit).Offset((page - 1) * limit).Find(&posts).Error
	
	return posts, total, err
}
