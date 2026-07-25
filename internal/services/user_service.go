package services

import (
	"errors"

	"gorm.io/gorm"

	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses/users"
	"easyrent-server/internal/repositories"
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

// GetMe retrieves the authenticated user's information based on the provided user ID.
func (s *UserService) GetMe(
	userID string,
) (*users.MeResponse, error) {
	user, err := s.userRepository.GetByID(userID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.RecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &users.MeResponse{
		ID:             user.ID,
		Email:          user.Account.Email,
		PhoneNumber:    user.Account.PhoneNumber,
		Role:           user.Account.Role,
		Status:         user.Status,
		EmailVerified:  user.EmailVerified,
		FullName:       user.UserInfo.FullName,
		AvatarURL:      user.UserInfo.AvatarURL,
		Gender:         user.UserInfo.Gender,
		Birthday:       user.UserInfo.Birthday,
		Address:        user.UserInfo.Address,
		IdentityNumber: user.UserInfo.IdentityNumber,
		Bio:            user.UserInfo.Bio,
		Occupation:     user.UserInfo.Occupation,
		LastLoginAt:    user.LastLoginAt,
	}, nil
}

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
	}

	return s.userRepository.Update(userID, updateData)
}
