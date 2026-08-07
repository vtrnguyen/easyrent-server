package services

import (
	"errors"
	"strings"

	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"

	"gorm.io/gorm"
)

type RentalRequestService struct {
	repository *repositories.RentalRequestRepository
	properties *repositories.PropertyRepository
}

func NewRentalRequestService() *RentalRequestService {
	return &RentalRequestService{repository: &repositories.RentalRequestRepository{}, properties: &repositories.PropertyRepository{}}
}

func (s *RentalRequestService) Create(tenantID, propertyID, message string) error {
	property, err := s.properties.GetByID(propertyID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}
	if property.Status != "available" || property.OwnerID == tenantID {
		return apperrors.Forbidden
	}
	request := &models.RentalRequest{TenantID: tenantID, PropertyID: propertyID, Message: strings.TrimSpace(message), Status: "pending"}
	return s.repository.Create(request)
}
