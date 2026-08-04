package repositories

import (
	"easyrent-server/internal/database"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/models"

	"gorm.io/gorm"
)

type PropertyRepository struct{}

// Create creates a property and its media in a single transaction.
func (r *PropertyRepository) Create(
	property *models.Property,
	images []models.PropertyImage,
	videos []models.PropertyVideo,
	utilities []models.Utility,
) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(property).Error; err != nil {
			return err
		}

		for i := range images {
			images[i].PropertyID = property.ID
			if err := tx.Create(&images[i]).Error; err != nil {
				return err
			}
		}

		for i := range videos {
			videos[i].PropertyID = property.ID
			if err := tx.Create(&videos[i]).Error; err != nil {
				return err
			}
		}

		if len(utilities) > 0 {
			propertyUtilities := make([]models.PropertyUtility, 0, len(utilities))
			for _, utility := range utilities {
				propertyUtilities = append(propertyUtilities, models.PropertyUtility{
					PropertyID: property.ID,
					UtilityID:  utility.ID,
				})
			}

			if err := tx.Create(&propertyUtilities).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetByID retrieves a property with its media and utilities.
func (r *PropertyRepository) GetByID(
	propertyID string,
) (*models.Property, error) {
	var property models.Property

	err := database.DB.
		Preload("Images").
		Preload("Videos").
		Preload("Utilities").
		First(&property, "id = ?", propertyID).
		Error
	if err != nil {
		return nil, err
	}

	return &property, nil
}

// Update updates a property and optionally replaces its media.
func (r *PropertyRepository) Update(
	propertyID string,
	data map[string]interface{},
	images []models.PropertyImage,
	videos []models.PropertyVideo,
	utilities []models.Utility,
) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Property{}).
			Where("id = ?", propertyID).
			Updates(data).
			Error; err != nil {
			return err
		}

		if images != nil {
			if err := tx.Where("property_id = ?", propertyID).
				Delete(&models.PropertyImage{}).
				Error; err != nil {
				return err
			}

			for i := range images {
				images[i].PropertyID = propertyID
				if err := tx.Create(&images[i]).Error; err != nil {
					return err
				}
			}
		}

		if videos != nil {
			if err := tx.Where("property_id = ?", propertyID).
				Delete(&models.PropertyVideo{}).
				Error; err != nil {
				return err
			}

			for i := range videos {
				videos[i].PropertyID = propertyID
				if err := tx.Create(&videos[i]).Error; err != nil {
					return err
				}
			}
		}

		if utilities != nil {
			if err := tx.Where("property_id = ?", propertyID).
				Delete(&models.PropertyUtility{}).
				Error; err != nil {
				return err
			}

			if len(utilities) > 0 {
				propertyUtilities := make([]models.PropertyUtility, 0, len(utilities))
				for _, utility := range utilities {
					propertyUtilities = append(propertyUtilities, models.PropertyUtility{
						PropertyID: propertyID,
						UtilityID:  utility.ID,
					})
				}

				if err := tx.Create(&propertyUtilities).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// Delete removes a property and all related records in a single transaction.
func (r *PropertyRepository) Delete(
	propertyID string,
) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("property_id = ?", propertyID).
			Delete(&models.PropertyUtility{}).
			Error; err != nil {
			return err
		}

		if err := tx.Where("property_id = ?", propertyID).
			Delete(&models.PropertyImage{}).
			Error; err != nil {
			return err
		}

		if err := tx.Where("property_id = ?", propertyID).
			Delete(&models.PropertyVideo{}).
			Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Property{}, "id = ?", propertyID).Error; err != nil {
			return err
		}

		return nil
	})
}

// Search retrieves a list of properties with filtering, sorting, and pagination.
func (r *PropertyRepository) Search(
	req requests.SearchRequest,
	ownerID string,
) ([]models.Property, int64, error) {
	var properties []models.Property
	var total int64

	fieldMap := map[string]string{
		"owner_id":            "properties.owner_id",
		"property_type_id":    "properties.property_type_id",
		"title":               "properties.title",
		"type":                "properties.type",
		"province":            "properties.province",
		"district":            "properties.district",
		"ward":                "properties.ward",
		"address":             "properties.address",
		"latitude":            "properties.latitude",
		"longitude":           "properties.longitude",
		"area":                "properties.area",
		"max_people":          "properties.max_people",
		"number_of_bedrooms":  "properties.number_of_bedrooms",
		"number_of_bathrooms": "properties.number_of_bathrooms",
		"price":               "properties.price",
		"electricity_price":   "properties.electricity_price",
		"water_price":         "properties.water_price",
		"status":              "properties.status",
		"created_at":          "properties.created_at",
		"updated_at":          "properties.updated_at",
	}

	query := database.DB.Model(&models.Property{})
	if ownerID != "" {
		query = query.Where("owner_id = ?", ownerID)
	}

	query = database.ApplyFilters(query, req.Filters, fieldMap, req.FilterLogic)
	query.Count(&total)

	query = database.ApplySorts(query, req.Sorts, fieldMap)

	err := query.
		Preload("Images").
		Preload("Videos").
		Preload("Utilities").
		Limit(req.Limit).
		Offset((req.Page - 1) * req.Limit).
		Find(&properties).
		Error

	return properties, total, err
}
