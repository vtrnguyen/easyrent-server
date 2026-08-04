package services

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"
)

type UtilityService struct {
	UtilityRepository *repositories.UtilityRepository
}

// NewUtilityService creates a new instance of UtilityService with the necessary dependencies.
func NewUtilityService() *UtilityService {
	return &UtilityService{
		UtilityRepository: &repositories.UtilityRepository{},
	}
}

// GetAll retrieves all utilities from the repository.
func (s *UtilityService) GetAll() ([]responses.UtilityResponse, error) {
	utilities, err := s.UtilityRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var utilityResponses []responses.UtilityResponse
	for _, utility := range utilities {
		utilityResponses = append(utilityResponses, responses.UtilityResponse{
			ID:           utility.ID,
			Code:         utility.Code,
			DisplayName:  utility.DisplayName,
			DisplayOrder: utility.DisplayOrder,
		})
	}

	return utilityResponses, nil
}

// Create creates a new utility using the repository.
func (s *UtilityService) Create(utility *requests.UtilityRequest) error {
	utilityData := &models.Utility{
		Code:         utility.Code,
		DisplayName:  utility.DisplayName,
		DisplayOrder: utility.DisplayOrder,
	}

	return s.UtilityRepository.Create(utilityData)
}

// Update updates an existing utility using the repository.
func (s *UtilityService) Update(utility *requests.UtilityRequest) error {
	utilityData := &models.Utility{
		Code:         utility.Code,
		DisplayName:  utility.DisplayName,
		DisplayOrder: utility.DisplayOrder,
	}

	return s.UtilityRepository.Update(utilityData)
}

// Delete deletes a utility by its ID using the repository.
func (s *UtilityService) Delete(utilityID string) error {
	return s.UtilityRepository.Delete(utilityID)
}
