package services

import (
	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/repositories"
	"easyrent-server/internal/utils"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

// NewUserService creates a new instance of UserService with the necessary dependencies.
func NewUserService() *UserService {
	return &UserService{
		userRepository: &repositories.UserRepository{},
	}
}

// GetByID retrieves a user's information based on the provided user ID.
func (s *UserService) GetByID(
	userID string,
) (*responses.UserResponse, error) {
	user, err := s.userRepository.GetByID(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.RecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &responses.UserResponse{
		ID:             user.ID,
		Email:          user.Account.Email,
		PhoneNumber:    user.Account.PhoneNumber,
		Role:           user.Account.Role,
		Status:         user.Status,
		EmailVerified:  user.EmailVerified,
		FullName:       user.FullName,
		AvatarURL:      user.AvatarURL,
		Gender:         user.Gender,
		Birthday:       user.Birthday,
		Address:        user.Address,
		IdentityNumber: user.IdentityNumber,
		Bio:            user.Bio,
		Occupation:     user.Occupation,
		LastLoginAt:    user.LastLoginAt,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}, nil
}

// UpdateMe updates the authenticated user's information based on the provided user ID and request data.
func (s *UserService) UpdateMe(
	userID string,
	req requests.UpdateMeRequest,
) error {
	updateData := map[string]interface{}{
		"full_name":       req.FullName,
		"avatar_url":      req.AvatarURL,
		"gender":          req.Gender,
		"birthday":        req.Birthday,
		"address":         req.Address,
		"identity_number": req.IdentityNumber,
		"bio":             req.Bio,
		"occupation":      req.Occupation,
		"updated_at":      gorm.Expr("NOW()"),
	}

	return s.userRepository.Update(userID, updateData)
}

// Search retrieves a paginated list of users based on the provided search criteria.
func (s *UserService) Search(
	req requests.SearchRequest,
) (*responses.PaginatedResponse[responses.UserResponse], error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	userList, total, err := s.userRepository.Search(req)

	if err != nil {
		return nil, err
	}

	items := make(
		[]responses.UserResponse,
		0,
		len(userList),
	)

	for _, user := range userList {
		items = append(
			items,
			responses.UserResponse{
				ID:             user.ID,
				Email:          user.Account.Email,
				PhoneNumber:    user.Account.PhoneNumber,
				Role:           user.Account.Role,
				Status:         user.Status,
				EmailVerified:  user.EmailVerified,
				FullName:       user.FullName,
				AvatarURL:      user.AvatarURL,
				Gender:         user.Gender,
				Birthday:       user.Birthday,
				Address:        user.Address,
				IdentityNumber: user.IdentityNumber,
				Bio:            user.Bio,
				Occupation:     user.Occupation,
				LastLoginAt:    user.LastLoginAt,
				CreatedAt:      user.CreatedAt,
				UpdatedAt:      user.UpdatedAt,
			},
		)
	}

	totalPages := utils.CalculateTotalPages(total, req.Limit)

	return &responses.PaginatedResponse[responses.UserResponse]{
		Items:      items,
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}, nil
}
