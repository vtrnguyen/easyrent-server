package services

import (
	"errors"

	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/repositories"
	"easyrent-server/internal/utils"

	"gorm.io/gorm"
)

type PostFavoriteService struct {
	repository  *repositories.PostFavoriteRepository
	postService *PostService
}

// NewPostFavoriteService creates a new PostFavoriteService instance.
func NewPostFavoriteService() *PostFavoriteService {
	return &PostFavoriteService{repository: &repositories.PostFavoriteRepository{}, postService: NewPostService()}
}

// Add adds a post to the user's favorites.
func (s *PostFavoriteService) Add(userID, postID string) error {
	post, err := s.postService.postRepository.GetByID(postID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}
	if post.Status != "published" {
		return apperrors.Forbidden
	}

	return s.repository.Add(userID, postID)
}

// Remove removes a post from the user's favorites.
func (s *PostFavoriteService) Remove(userID, postID string) error {
	return s.repository.Remove(userID, postID)
}

// IDs retrieves the IDs of the user's favorite posts.
func (s *PostFavoriteService) IDs(userID string) ([]string, error) { return s.repository.IDs(userID) }

// Search retrieves the user's favorite posts with pagination.
func (s *PostFavoriteService) Search(userID string, page, limit int) (*responses.PaginatedResponse[responses.PostResponse], error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 12
	}

	posts, total, err := s.repository.Search(userID, page, limit)
	if err != nil {
		return nil, err
	}

	items := make([]responses.PostResponse, 0, len(posts))
	for _, post := range posts {
		items = append(items, s.postService.mapResponse(post))
	}
	
	return &responses.PaginatedResponse[responses.PostResponse]{Items: items, Total: total, Page: page, Limit: limit, TotalPages: utils.CalculateTotalPages(total, limit)}, nil
}
