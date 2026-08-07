package services

import (
	"errors"
	"strings"

	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"
	"easyrent-server/internal/utils"

	"gorm.io/gorm"
)

type PostInteractionService struct {
	repository *repositories.PostInteractionRepository
	posts      *repositories.PostRepository
}

func NewPostInteractionService() *PostInteractionService {
	return &PostInteractionService{repository: &repositories.PostInteractionRepository{}, posts: &repositories.PostRepository{}}
}

func (s *PostInteractionService) ensurePublished(postID string) error {
	post, err := s.posts.GetByID(postID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}
	if post.Status != "published" {
		return apperrors.Forbidden
	}
	return nil
}

func (s *PostInteractionService) Social(userID, postID string) (*responses.PostSocialResponse, error) {
	if err := s.ensurePublished(postID); err != nil {
		return nil, err
	}
	liked, count, err := s.repository.Social(userID, postID)
	if err != nil {
		return nil, err
	}
	return &responses.PostSocialResponse{Liked: liked, LikeCount: count}, nil
}

func (s *PostInteractionService) Like(userID, postID string) error {
	if err := s.ensurePublished(postID); err != nil {
		return err
	}
	return s.repository.Like(userID, postID)
}

func (s *PostInteractionService) Unlike(userID, postID string) error {
	return s.repository.Unlike(userID, postID)
}

func (s *PostInteractionService) Comments(postID string, page, limit int) (*responses.PaginatedResponse[responses.PostCommentResponse], error) {
	if err := s.ensurePublished(postID); err != nil {
		return nil, err
	}
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	comments, total, err := s.repository.Comments(postID, page, limit)
	if err != nil {
		return nil, err
	}
	items := buildCommentTree(comments)
	return &responses.PaginatedResponse[responses.PostCommentResponse]{Items: items, Total: total, Page: page, Limit: limit, TotalPages: utils.CalculateTotalPages(total, limit)}, nil
}

func (s *PostInteractionService) Comment(userID, postID, content string, parentCommentID *string) (*responses.PostCommentResponse, error) {
	if err := s.ensurePublished(postID); err != nil {
		return nil, err
	}
	if parentCommentID != nil {
		parent, err := s.repository.GetComment(*parentCommentID)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.RecordNotFound
		}
		if err != nil {
			return nil, err
		}
		if parent.PostID != postID {
			return nil, apperrors.Forbidden
		}
	}
	comment := &models.PostComment{UserID: userID, PostID: postID, ParentCommentID: parentCommentID, Content: strings.TrimSpace(content)}
	if err := s.repository.CreateComment(comment); err != nil {
		return nil, err
	}
	response := mapComment(*comment)
	return &response, nil
}

func mapComment(comment models.PostComment) responses.PostCommentResponse {
	return responses.PostCommentResponse{ID: comment.ID, UserID: comment.UserID, UserName: comment.User.FullName, UserAvatarURL: comment.User.AvatarURL, ParentCommentID: comment.ParentCommentID, Content: comment.Content, CreatedAt: comment.CreatedAt, Children: []responses.PostCommentResponse{}}
}

func buildCommentTree(comments []models.PostComment) []responses.PostCommentResponse {
	childrenByParent := make(map[string][]models.PostComment)
	var roots []models.PostComment
	for _, comment := range comments {
		if comment.ParentCommentID == nil {
			roots = append(roots, comment)
		} else {
			childrenByParent[*comment.ParentCommentID] = append(childrenByParent[*comment.ParentCommentID], comment)
		}
	}
	var build func(models.PostComment) responses.PostCommentResponse
	build = func(comment models.PostComment) responses.PostCommentResponse {
		response := mapComment(comment)
		for _, child := range childrenByParent[comment.ID] {
			response.Children = append(response.Children, build(child))
		}
		return response
	}
	items := make([]responses.PostCommentResponse, 0, len(roots))
	for _, root := range roots {
		items = append(items, build(root))
	}
	return items
}
