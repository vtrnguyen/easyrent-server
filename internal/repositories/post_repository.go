package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/models"

	"gorm.io/gorm"
)

type PostRepository struct{}

func (r *PostRepository) Create(post *models.Post) error {
	return database.DB.Create(post).Error
}

func (r *PostRepository) GetByID(id string) (*models.Post, error) {
	var post models.Post
	err := database.DB.Preload("Property").First(&post, "id = ?", id).Error
	return &post, err
}

func (r *PostRepository) Update(id string, data map[string]interface{}) error {
	return database.DB.Model(&models.Post{}).Where("id = ?", id).Updates(data).Error
}

func (r *PostRepository) Delete(id string) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		for _, model := range []interface{}{&models.PostLike{}, &models.PostFavorite{}, &models.PostComment{}, &models.PostTagRelation{}} {
			if err := tx.Where("post_id = ?", id).Delete(model).Error; err != nil {
				return err
			}
		}
		return tx.Delete(&models.Post{}, "id = ?", id).Error
	})
}

func (r *PostRepository) Search(req requests.SearchRequest, authorID string) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64
	fieldMap := map[string]string{
		"property_id": "posts.property_id", "title": "posts.title",
		"status": "posts.status", "published_at": "posts.published_at", "created_at": "posts.created_at",
	}
	query := database.DB.Model(&models.Post{}).Where("author_id = ?", authorID)
	query = database.ApplyFilters(query, req.Filters, fieldMap, req.FilterLogic)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	query = database.ApplySorts(query, req.Sorts, fieldMap)
	err := query.Preload("Property").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&posts).Error
	return posts, total, err
}
