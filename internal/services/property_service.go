package services

import (
	"easyrent-server/internal/apperrors"
	"easyrent-server/internal/constants"
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/dto/responses"
	"easyrent-server/internal/models"
	"easyrent-server/internal/repositories"
	"easyrent-server/internal/utils"
	"errors"
	"mime/multipart"

	"gorm.io/gorm"
)

type PropertyService struct {
	propertyRepository *repositories.PropertyRepository
	fileService        *FileService
}

// NewPropertyService creates a new instance of PropertyService.
func NewPropertyService() *PropertyService {
	return &PropertyService{
		propertyRepository: &repositories.PropertyRepository{},
		fileService:        NewFileService(),
	}
}

func (s *PropertyService) buildImages(
	files []*multipart.FileHeader,
) ([]models.PropertyImage, []string, error) {
	if len(files) == 0 {
		return nil, nil, nil
	}

	images := make([]models.PropertyImage, 0, len(files))
	urls := make([]string, 0, len(files))

	for index, file := range files {
		url, err := s.fileService.SavePropertyImage(file)
		if err != nil {
			return nil, urls, err
		}

		urls = append(urls, url)
		images = append(images, models.PropertyImage{
			ImageURL:     url,
			IsThumbnail:  index == 0,
			DisplayOrder: index + 1,
		})
	}

	return images, urls, nil
}

func (s *PropertyService) buildVideos(
	files []*multipart.FileHeader,
) ([]models.PropertyVideo, []string, error) {
	if len(files) == 0 {
		return nil, nil, nil
	}

	videos := make([]models.PropertyVideo, 0, len(files))
	urls := make([]string, 0, len(files))

	for _, file := range files {
		url, err := s.fileService.SavePropertyVideo(file)
		if err != nil {
			return nil, urls, err
		}

		urls = append(urls, url)
		videos = append(videos, models.PropertyVideo{
			VideoURL: url,
		})
	}

	return videos, urls, nil
}

func (s *PropertyService) cleanupFiles(urls []string) {
	for _, url := range urls {
		_ = s.fileService.DeleteByURL(url)
	}
}

func (s *PropertyService) mapResponse(property models.Property) responses.PropertyResponse {
	images := make([]responses.PropertyImageResponse, 0, len(property.Images))
	for _, image := range property.Images {
		images = append(images, responses.PropertyImageResponse{
			ID:           image.ID,
			ImageURL:     image.ImageURL,
			IsThumbnail:  image.IsThumbnail,
			DisplayOrder: image.DisplayOrder,
		})
	}

	videos := make([]responses.PropertyVideoResponse, 0, len(property.Videos))
	for _, video := range property.Videos {
		videos = append(videos, responses.PropertyVideoResponse{
			ID:       video.ID,
			VideoURL: video.VideoURL,
		})
	}

	utilities := make([]string, 0, len(property.Utilities))
	for _, utility := range property.Utilities {
		utilities = append(utilities, utility.Code)
	}

	return responses.PropertyResponse{
		ID:                property.ID,
		OwnerID:           property.OwnerID,
		Title:             property.Title,
		Type:              property.Type,
		Description:       property.Description,
		Province:          property.Province,
		District:          property.District,
		Ward:              property.Ward,
		Address:           property.Address,
		Latitude:          property.Latitude,
		Longitude:         property.Longitude,
		Area:              property.Area,
		MaxPeople:         property.MaxPeople,
		NumberOfBedrooms:  property.NumberOfBedrooms,
		NumberOfBathrooms: property.NumberOfBathrooms,
		ExtraRoomInfos:    property.ExtraRoomInfos,
		Price:             property.Price,
		ElectricityPrice:  property.ElectricityPrice,
		WaterPrice:        property.WaterPrice,
		Status:            property.Status,
		Images:            images,
		Videos:            videos,
		Utilities:         utilities,
		CreatedAt:         property.CreatedAt,
		UpdatedAt:         property.UpdatedAt,
	}
}

// GetByID retrieves a property detail and enforces ownership for landlord users.
func (s *PropertyService) GetByID(
	actorID string,
	actorRole string,
	propertyID string,
) (*responses.PropertyResponse, error) {
	property, err := s.propertyRepository.GetByID(propertyID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperrors.RecordNotFound
	}
	if err != nil {
		return nil, err
	}

	if actorRole != string(constants.AccountRoleAdmin) && property.OwnerID != actorID {
		return nil, apperrors.Forbidden
	}

	response := s.mapResponse(*property)
	return &response, nil
}

// Create creates a new property and saves its media files.
func (s *PropertyService) Create(
	actorID string,
	actorRole string,
	req requests.PropertyRequest,
	imageFiles []*multipart.FileHeader,
	videoFiles []*multipart.FileHeader,
) error {
	ownerID := actorID
	if actorRole == string(constants.AccountRoleAdmin) && req.OwnerID != "" {
		ownerID = req.OwnerID
	}

	images, imageURLs, err := s.buildImages(imageFiles)
	if err != nil {
		s.cleanupFiles(imageURLs)
		return err
	}

	videos, videoURLs, err := s.buildVideos(videoFiles)
	if err != nil {
		s.cleanupFiles(imageURLs)
		s.cleanupFiles(videoURLs)
		return err
	}

	property := models.Property{
		OwnerID:           ownerID,
		Title:             req.Title,
		Type:              req.Type,
		Description:       req.Description,
		Province:          req.Province,
		District:          req.District,
		Ward:              req.Ward,
		Address:           req.Address,
		Latitude:          req.Latitude,
		Longitude:         req.Longitude,
		Area:              req.Area,
		MaxPeople:         req.MaxPeople,
		NumberOfBedrooms:  req.NumberOfBedrooms,
		NumberOfBathrooms: req.NumberOfBathrooms,
		ExtraRoomInfos:    req.ExtraRoomInfos,
		Price:             req.Price,
		ElectricityPrice:  req.ElectricityPrice,
		WaterPrice:        req.WaterPrice,
		Status:            req.Status,
	}

	err = s.propertyRepository.Create(&property, images, videos)
	if err != nil {
		s.cleanupFiles(imageURLs)
		s.cleanupFiles(videoURLs)
		return err
	}

	return nil
}

// Update updates a property and optionally replaces its media files.
func (s *PropertyService) Update(
	actorID string,
	actorRole string,
	propertyID string,
	req requests.PropertyRequest,
	imageFiles []*multipart.FileHeader,
	videoFiles []*multipart.FileHeader,
) error {
	property, err := s.propertyRepository.GetByID(propertyID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}

	if actorRole != string(constants.AccountRoleAdmin) && property.OwnerID != actorID {
		return apperrors.Forbidden
	}

	ownerID := property.OwnerID
	if actorRole == string(constants.AccountRoleAdmin) && req.OwnerID != "" {
		ownerID = req.OwnerID
	}

	images, imageURLs, err := s.buildImages(imageFiles)
	if err != nil {
		s.cleanupFiles(imageURLs)
		return err
	}

	videos, videoURLs, err := s.buildVideos(videoFiles)
	if err != nil {
		s.cleanupFiles(imageURLs)
		s.cleanupFiles(videoURLs)
		return err
	}

	updateData := map[string]interface{}{
		"owner_id":            ownerID,
		"title":               req.Title,
		"type":                req.Type,
		"description":         req.Description,
		"province":            req.Province,
		"district":            req.District,
		"ward":                req.Ward,
		"address":             req.Address,
		"latitude":            req.Latitude,
		"longitude":           req.Longitude,
		"area":                req.Area,
		"max_people":          req.MaxPeople,
		"number_of_bedrooms":  req.NumberOfBedrooms,
		"number_of_bathrooms": req.NumberOfBathrooms,
		"extra_room_infos":    req.ExtraRoomInfos,
		"price":               req.Price,
		"electricity_price":   req.ElectricityPrice,
		"water_price":         req.WaterPrice,
		"status":              req.Status,
		"updated_at":          gorm.Expr("NOW()"),
	}

	var oldImageURLs []string
	if len(imageFiles) > 0 {
		oldImageURLs = make([]string, 0, len(property.Images))
		for _, image := range property.Images {
			oldImageURLs = append(oldImageURLs, image.ImageURL)
		}
	}

	var oldVideoURLs []string
	if len(videoFiles) > 0 {
		oldVideoURLs = make([]string, 0, len(property.Videos))
		for _, video := range property.Videos {
			oldVideoURLs = append(oldVideoURLs, video.VideoURL)
		}
	}

	err = s.propertyRepository.Update(propertyID, updateData, images, videos)
	if err != nil {
		s.cleanupFiles(imageURLs)
		s.cleanupFiles(videoURLs)
		return err
	}

	s.cleanupFiles(oldImageURLs)
	s.cleanupFiles(oldVideoURLs)

	return nil
}

// Delete removes a property and its stored media files.
func (s *PropertyService) Delete(
	actorID string,
	actorRole string,
	propertyID string,
) error {
	property, err := s.propertyRepository.GetByID(propertyID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperrors.RecordNotFound
	}
	if err != nil {
		return err
	}

	if actorRole != string(constants.AccountRoleAdmin) && property.OwnerID != actorID {
		return apperrors.Forbidden
	}

	imageURLs := make([]string, 0, len(property.Images))
	for _, image := range property.Images {
		imageURLs = append(imageURLs, image.ImageURL)
	}

	videoURLs := make([]string, 0, len(property.Videos))
	for _, video := range property.Videos {
		videoURLs = append(videoURLs, video.VideoURL)
	}

	if err := s.propertyRepository.Delete(propertyID); err != nil {
		return err
	}

	s.cleanupFiles(imageURLs)
	s.cleanupFiles(videoURLs)

	return nil
}

// Search retrieves properties with pagination, filtering, and sorting.
func (s *PropertyService) Search(
	actorID string,
	actorRole string,
	req requests.SearchRequest,
) (*responses.PaginatedResponse[responses.PropertyResponse], error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}

	ownerID := ""
	if actorRole != string(constants.AccountRoleAdmin) {
		ownerID = actorID
	}

	properties, total, err := s.propertyRepository.Search(req, ownerID)
	if err != nil {
		return nil, err
	}

	items := make([]responses.PropertyResponse, 0, len(properties))
	for _, property := range properties {
		items = append(items, s.mapResponse(property))
	}

	return &responses.PaginatedResponse[responses.PropertyResponse]{
		Items:      items,
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: utils.CalculateTotalPages(total, req.Limit),
	}, nil
}
