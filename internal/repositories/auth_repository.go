package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/models"
	"time"

	"gorm.io/gorm"
)

type AuthRepository struct{}

// CreateUser creates a new user, account, and user info in the database within a transaction.
func (r *AuthRepository) CreateUser(
	tx *gorm.DB,
	user *models.User,
	account *models.Account,
) error {
	if err := tx.Create(user).Error; err != nil {
		return err
	}

	account.UserID = user.ID

	if err := tx.Create(account).Error; err != nil {
		return err
	}

	return nil
}

// IsEmailExists checks if an email already exists in the accounts table.
func (r *AuthRepository) IsEmailExists(email string, excludeUserID ...string) (bool, error) {
	var count int64

	query := database.DB.
		Model(&models.Account{}).
		Where("email = ?", email)

	if len(excludeUserID) > 0 {
		query = query.Where("user_id != ?", excludeUserID[0])
	}

	err := query.Count(&count).Error

	return count > 0, err
}

// IsPhoneNumberExists checks if a phone number already exists in the accounts table.
func (r *AuthRepository) IsPhoneNumberExists(phoneNumber string, excludeUserID ...string) (bool, error) {
	var count int64

	query := database.DB.
		Model(&models.Account{}).
		Where("phone_number = ?", phoneNumber)

	if len(excludeUserID) > 0 {
		query = query.Where("user_id != ?", excludeUserID[0])
	}

	err := query.Count(&count).Error

	return count > 0, err
}

// FindUserByLogin retrieves an account by email or phone number.
func (r *AuthRepository) FindUserByLogin(
	login string,
) (*models.User, error) {
	var user models.User

	err := database.DB.
		Preload("Account").
		Where("id IN (?)",
			database.DB.
				Table("accounts").
				Select("user_id").
				Where(
					"email = ? OR phone_number = ?", login, login,
				),
		).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateLastLoginAt updates the last login timestamp for a user.
func (r *AuthRepository) UpdateLastLoginAt(userID string) error {
	now := time.Now()
	return database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("last_login_at", &now).
		Error
}

// BeginTransaction starts a new database transaction and returns the transaction object.
func (r *AuthRepository) BeginTransaction() *gorm.DB {
	return database.DB.Begin()
}
