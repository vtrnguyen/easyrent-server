package services

import (
	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/constants"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"
	"easyrent-server/internal/utils"
)

type AuthService struct {
	authRepository *repositories.AuthRepository
}

// NewAuthService creates a new instance of AuthService with the necessary dependencies.
func NewAuthService() *AuthService {
	return &AuthService{
		authRepository: &repositories.AuthRepository{},
	}
}

// Register handles the user registration process. It checks for existing email and phone number,
func (s *AuthService) Register(
	req requests.RegisterRequest,
) (responses.AuthResponse, error) {
	isEmailExists, err := s.authRepository.IsEmailExists(req.Email)
	if err != nil {
		return responses.AuthResponse{}, err
	}
	if isEmailExists {
		return responses.AuthResponse{}, apperrors.EmailAlreadyExists
	}

	isPhoneNumberExists, err := s.authRepository.IsPhoneNumberExists(req.PhoneNumber)
	if err != nil {
		return responses.AuthResponse{}, err
	}
	if isPhoneNumberExists {
		return responses.AuthResponse{}, apperrors.PhoneAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	user := &models.User{
		FullName: req.FullName,
		Gender:   req.Gender,
		Status:   string(constants.UserStatusActive),
	}

	account := &models.Account{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    hashedPassword,
		Role:        string(constants.AccountRoleTenant),
	}

	tx := s.authRepository.BeginTransaction()

	if err := s.authRepository.CreateUser(
		tx,
		user,
		account,
	); err != nil {
		tx.Rollback()
		return responses.AuthResponse{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return responses.AuthResponse{}, err
	}

	accessToken, err := utils.GenerateAccessToken(
		user.ID,
		account.Role,
	)

	return responses.AuthResponse{
		AccessToken: accessToken,
		UserID:      user.ID,
		FullName:    user.FullName,
		AvatarURL:   user.AvatarURL,
		Role:        account.Role,
		Status:      user.Status,
		LastLoginAt: user.LastLoginAt,
	}, nil
}

// Login handles the user login process. It verifies the provided email or phone number and password,
func (s *AuthService) Login(
	req requests.LoginRequest,
) (responses.AuthResponse, error) {
	user, err := s.authRepository.
		FindUserByLogin(req.Identifier)
	if err != nil {
		return responses.AuthResponse{}, apperrors.InvalidLoginCredentials
	}

	isPasswordValid := utils.ComparePassword(
		req.Password,
		user.Account.Password,
	)
	if !isPasswordValid {
		return responses.AuthResponse{}, apperrors.InvalidLoginCredentials
	}

	if err := s.authRepository.UpdateLastLoginAt(user.ID); err != nil {
		return responses.AuthResponse{}, err
	}

	accessToken, err := utils.GenerateAccessToken(
		user.ID,
		user.Account.Role,
	)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.AuthResponse{
		AccessToken: accessToken,
		UserID:      user.ID,
		FullName:    user.FullName,
		AvatarURL:   user.AvatarURL,
		Role:        user.Account.Role,
		Status:      user.Status,
		LastLoginAt: user.LastLoginAt,
	}, nil
}
