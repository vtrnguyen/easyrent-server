package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/models"
)

type RentalRequestRepository struct{}

func (r *RentalRequestRepository) Create(request *models.RentalRequest) error {
	return database.DB.Create(request).Error
}
