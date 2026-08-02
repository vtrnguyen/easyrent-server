package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct{}

// Create creates a new user with account information using a database transaction.
func (r *UserRepository) Create(
	user *models.User,
	account *models.Account,
) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		account.UserID = user.ID

		if err := tx.Create(account).Error; err != nil {
			return err
		}

		return nil
	})
}

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
func (r *UserRepository) UpdateMe(
	userID string,
	data map[string]interface{},
) error {
	return database.DB.
		Model(&models.User{}).
		Where("id = ?", userID).
		Updates(data).
		Error
}

// UpdateProfile updates both user and account data in a single transaction.
func (r *UserRepository) UpdateProfile(
	userID string,
	userData map[string]interface{},
	accountData map[string]interface{},
) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if len(userData) > 0 {
			if err := tx.Model(&models.User{}).
				Where("id = ?", userID).
				Updates(userData).
				Error; err != nil {
				return err
			}
		}

		if len(accountData) > 0 {
			if err := tx.Model(&models.Account{}).
				Where("user_id = ?", userID).
				Updates(accountData).
				Error; err != nil {
				return err
			}
		}

		return nil
	})
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
		"last_login_at":   "users.last_login_at",
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
