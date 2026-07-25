package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/models"
)

type UserRepository struct{}

// GetByID retrieves a user by their ID, including their associated account and user info.
func (r *UserRepository) GetByID(
	userID string,
) (*models.User, error) {
	var user models.User

	err := database.DB.
		Preload("Account").
		First(&user, "id = ?", userID).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates the user information for a given user ID with the provided data.
func (r *UserRepository) Update(
	userID string,
	data map[string]interface{},
) error {
	return database.DB.
		Model(&models.User{}).
		Where("user_id = ?", userID).
		Updates(data).
		Error
}
