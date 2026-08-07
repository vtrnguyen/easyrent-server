package services

import (
	"errors"
	"strings"
	"time"

	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/constants"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"
	"easyrent-server/internal/utils"
	"gorm.io/gorm"
)

type PostService struct {
	postRepository     *repositories.PostRepository
	propertyRepository *repositories.PropertyRepository
}

func NewPostService() *PostService {
	return &PostService{postRepository: &repositories.PostRepository{}, propertyRepository: &repositories.PropertyRepository{}}
}

func (s *PostService) mapResponse(post models.Post) responses.PostResponse {
	post.ContentType = "plain_text"
	content := post.PlainContent
	if post.MarkdownContent != "" {
		post.ContentType = "markdown"
		content = post.MarkdownContent
	}
	thumbnailURL := ""
	for _, image := range post.Property.Images {
		if image.IsThumbnail {
			thumbnailURL = image.ImageURL
			break
		}
	}
	if thumbnailURL == "" && len(post.Property.Images) > 0 {
		thumbnailURL = post.Property.Images[0].ImageURL
	}
	address := strings.Join(filterNonEmpty([]string{post.Property.Address, post.Property.Ward, post.Property.District, post.Property.Province}), ", ")
	return responses.PostResponse{ID: post.ID, PropertyID: post.PropertyID, AuthorID: post.AuthorID, Title: post.Title,
		ContentType: post.ContentType, Content: content, Status: post.Status, PublishedAt: post.PublishedAt,
		ExpiresAt: post.ExpiresAt, PropertyTitle: post.Property.Title, PropertyPrice: post.Property.Price, PropertyArea: post.Property.Area, PropertyAddress: address, ThumbnailURL: thumbnailURL, CreatedAt: post.CreatedAt}
}

func filterNonEmpty(values []string) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

func (s *PostService) ensurePropertyOwner(propertyID, actorID string) error {
	property, err := s.propertyRepository.GetByID(propertyID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}
	if property.OwnerID != actorID {
		return apperrors.Forbidden
	}
	return nil
}

func (s *PostService) SearchPublished(req requests.SearchRequest) (*responses.PaginatedResponse[responses.PostResponse], error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	posts, total, err := s.postRepository.SearchPublished(req)
	if err != nil {
		return nil, err
	}
	items := make([]responses.PostResponse, 0, len(posts))
	for _, post := range posts {
		items = append(items, s.mapResponse(post))
	}
	return &responses.PaginatedResponse[responses.PostResponse]{Items: items, Total: total, Page: req.Page, Limit: req.Limit, TotalPages: utils.CalculateTotalPages(total, req.Limit)}, nil
}

func contentFields(contentType, content string) (string, string) {
	if contentType == "markdown" {
		return content, ""
	}
	return "", content
}

func (s *PostService) Create(actorID string, req requests.PostRequest) error {
	if err := s.ensurePropertyOwner(req.PropertyID, actorID); err != nil {
		return err
	}
	markdown, plain := contentFields(req.ContentType, req.Content)
	post := models.Post{PropertyID: req.PropertyID, AuthorID: actorID, Title: req.Title, ContentType: req.ContentType,
		MarkdownContent: markdown, PlainContent: plain, Status: req.Status}
	if req.Status == "published" {
		now := time.Now()
		post.PublishedAt = &now
	}
	return s.postRepository.Create(&post)
}

func (s *PostService) GetByID(actorID, actorRole, id string) (*responses.PostResponse, error) {
	post, err := s.postRepository.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.RecordNotFound
	}
	if err != nil {
		return nil, err
	}
	if actorRole == string(constants.AccountRoleLandlord) && post.AuthorID != actorID {
		return nil, apperrors.Forbidden
	}
	if actorRole == string(constants.AccountRoleTenant) && post.Status != "published" {
		return nil, apperrors.Forbidden
	}
	result := s.mapResponse(*post)
	return &result, nil
}

func (s *PostService) Update(actorID, id string, req requests.PostRequest) error {
	post, err := s.postRepository.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}
	if post.AuthorID != actorID {
		return apperrors.Forbidden
	}
	if err := s.ensurePropertyOwner(req.PropertyID, actorID); err != nil {
		return err
	}
	markdown, plain := contentFields(req.ContentType, req.Content)
	data := map[string]interface{}{"property_id": req.PropertyID, "title": req.Title,
		"markdown_content": markdown, "plain_content": plain, "status": req.Status}
	if req.Status == "published" && post.PublishedAt == nil {
		data["published_at"] = time.Now()
	}
	return s.postRepository.Update(id, data)
}

func (s *PostService) Delete(actorID, id string) error {
	post, err := s.postRepository.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}
	if post.AuthorID != actorID {
		return apperrors.Forbidden
	}
	return s.postRepository.Delete(id)
}

func (s *PostService) Search(actorID string, req requests.SearchRequest) (*responses.PaginatedResponse[responses.PostResponse], error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}
	posts, total, err := s.postRepository.Search(req, actorID)
	if err != nil {
		return nil, err
	}
	items := make([]responses.PostResponse, 0, len(posts))
	for _, post := range posts {
		items = append(items, s.mapResponse(post))
	}
	return &responses.PaginatedResponse[responses.PostResponse]{Items: items, Total: total, Page: req.Page, Limit: req.Limit,
		TotalPages: utils.CalculateTotalPages(total, req.Limit)}, nil
}
