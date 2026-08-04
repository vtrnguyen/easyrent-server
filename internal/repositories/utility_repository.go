package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/models"

	"gorm.io/gorm"
)

type UtilityRepository struct{}

// GetAll retrieves all utilities from the database.
func (r *UtilityRepository) GetAll() ([]models.Utility, error) {
	var utilities []models.Utility

	err := database.DB.Find(&utilities).Error
	if err != nil {
		return nil, err
	}

	return utilities, nil
}

// FindByIDs retrieves utilities by their IDs.
func (r *UtilityRepository) FindByIDs(utilityIDs []string) ([]models.Utility, error) {
	if len(utilityIDs) == 0 {
		return []models.Utility{}, nil
	}

	var utilities []models.Utility
	err := database.DB.Where("id IN ?", utilityIDs).Find(&utilities).Error
	if err != nil {
		return nil, err
	}

	return utilities, nil
}

// Create creates a new utility in the database within a transaction.
func (r *UtilityRepository) Create(utility *models.Utility) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(utility).Error; err != nil {
			return err
		}

		return nil
	})
}

// Update updates an existing utility in the database within a transaction.
func (r *UtilityRepository) Update(utility *models.Utility) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(utility).Error; err != nil {
			return err
		}

		return nil
	})
}

// Delete deletes a utility from the database within a transaction.
func (r *UtilityRepository) Delete(utilityID string) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Utility{}, "id = ?", utilityID).Error; err != nil {
			return err
		}

		return nil
	})
}
