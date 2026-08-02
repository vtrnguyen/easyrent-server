package services

import (
	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"
	email "easyrent-server/internal/shared/job"
	"easyrent-server/internal/utils"
	"errors"
	"fmt"
	"mime/multipart"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository *repositories.UserRepository
	authRepository *repositories.AuthRepository
	emailService   *EmailService
	fileService    *FileService
}

// NewUserService creates a new instance of UserService with the necessary dependencies.
func NewUserService() *UserService {
	emailService := NewEmailService()
	fileService := NewFileService()

	return &UserService{
		userRepository: &repositories.UserRepository{},
		authRepository: &repositories.AuthRepository{},
		emailService:   emailService,
		fileService:    fileService,
	}
}

// Create creates a new user based on the provided request data. It checks for existing email, hashes the password, and sends a welcome email to the user.
func (s *UserService) Create(
	req requests.CreateUserRequest,
	avatar *multipart.FileHeader,
) error {
	avatarURL := ""

	if avatar != nil {
		url, err := s.fileService.SaveAvatar(avatar)

		if err != nil {
			return err
		}

		avatarURL = url
	}

	exists, err := s.authRepository.IsEmailExists(req.Email)

	if err != nil {
		return err
	}

	if exists {
		return apperrors.EmailAlreadyExists
	}

	password := utils.GeneratePassword(6)

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	birthday, err := utils.ParseDate(req.Birthday)

	if err != nil {
		return err
	}

	fmt.Printf("Email: %s", req.Email)

	user := models.User{
		Status:         req.Status,
		FullName:       req.FullName,
		Gender:         req.Gender,
		Birthday:       birthday,
		Occupation:     req.Occupation,
		IdentityNumber: req.IdentityNumber,
		Address:        req.Address,
		Bio:            req.Bio,
		AvatarURL:      avatarURL,
	}

	account := models.Account{
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    hashedPassword,
		Role:        req.Role,
	}

	err = s.userRepository.Create(
		&user,
		&account,
	)

	if err != nil {
		return err
	}

	email.Queue <- email.Job{
		To:       req.Email,
		Subject:  "Welcome to EasyRent",
		Template: "welcome.html",
		Data: map[string]string{
			"Name":     req.FullName,
			"Email":    req.Email,
			"Password": password,
			"LoginURL": "http://localhost:3000/auth/login",
		},
	}

	return nil
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
