package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/models"
)

type PostInteractionRepository struct{}

func (r *PostInteractionRepository) Social(userID, postID string) (bool, int64, error) {
	var likeCount int64
	if err := database.DB.Model(&models.PostLike{}).Where("post_id = ?", postID).Count(&likeCount).Error; err != nil {
		return false, 0, err
	}
	var userLikeCount int64
	if err := database.DB.Model(&models.PostLike{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&userLikeCount).Error; err != nil {
		return false, 0, err
	}
	return userLikeCount > 0, likeCount, nil
}

func (r *PostInteractionRepository) Like(userID, postID string) error {
	like := models.PostLike{UserID: userID, PostID: postID}
	return database.DB.FirstOrCreate(&like, "user_id = ? AND post_id = ?", userID, postID).Error
}

func (r *PostInteractionRepository) Unlike(userID, postID string) error {
	return database.DB.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&models.PostLike{}).Error
}

func (r *PostInteractionRepository) Comments(postID string, page, limit int) ([]models.PostComment, int64, error) {
	var roots []models.PostComment
	var total int64
	query := database.DB.Model(&models.PostComment{}).Where("post_id = ? AND parent_comment_id IS NULL", postID)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Preload("User").Order("created_at DESC").Limit(limit).Offset((page - 1) * limit).Find(&roots).Error; err != nil {
		return nil, 0, err
	}

	comments := append([]models.PostComment{}, roots...)
	parentIDs := make([]string, 0, len(roots))
	for _, comment := range roots {
		parentIDs = append(parentIDs, comment.ID)
	}
	for len(parentIDs) > 0 {
		var children []models.PostComment
		if err := database.DB.Preload("User").Where("post_id = ? AND parent_comment_id IN ?", postID, parentIDs).Order("created_at ASC").Find(&children).Error; err != nil {
			return nil, 0, err
		}
		comments = append(comments, children...)
		parentIDs = parentIDs[:0]
		for _, child := range children {
			parentIDs = append(parentIDs, child.ID)
		}
	}
	return comments, total, nil
}

func (r *PostInteractionRepository) CreateComment(comment *models.PostComment) error {
	if err := database.DB.Create(comment).Error; err != nil {
		return err
	}
	return database.DB.Preload("User").First(comment, "id = ?", comment.ID).Error
}

func (r *PostInteractionRepository) GetComment(id string) (*models.PostComment, error) {
	var comment models.PostComment
	err := database.DB.First(&comment, "id = ?", id).Error
	return &comment, err
}
