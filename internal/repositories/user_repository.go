package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/dto/requests"
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

// Search retrieves a list of users based on the provided search request, including filtering, sorting, and pagination.
func (r *UserRepository) Search(
	req requests.SearchRequest,
) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	fieldMap := map[string]string{
		"full_name":       "users.full_name",
		"email":           "accounts.email",
		"phone_number":    "accounts.phone_number",
		"role":            "accounts.role",
		"status":          "users.status",
		"gender":          "users.gender",
		"email_verified":  "users.email_verified",
		"address":         "users.address",
		"identity_number": "users.identity_number",
		"occupation":      "users.occupation",
		"created_at":      "users.created_at",
	}

	query := database.DB.
		Debug().
		Model(&models.User{}).
		Joins("JOIN accounts ON accounts.user_id = users.id")

	query = database.ApplyFilters(query, req.Filters, fieldMap, req.FilterLogic)

	query.Count(&total)

	query = database.ApplySorts(
		query,
		req.Sorts,
		fieldMap,
	)

	err := query.
		Preload("Account").
		Limit(req.Limit).
		Offset((req.Page - 1) * req.Limit).
		Find(&users).
		Error

	return users, total, err
}
