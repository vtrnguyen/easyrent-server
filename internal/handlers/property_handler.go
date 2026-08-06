package handlers

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	propertyService *services.PropertyService
}

// NewPropertyHandler creates a new instance of PropertyHandler with the necessary dependencies.
func NewPropertyHandler() *PropertyHandler {
	return &PropertyHandler{
		propertyService: services.NewPropertyService(),
	}
}

// Create handles property creation for admin and landlord roles.
func (h *PropertyHandler) Create(c *gin.Context) {
	var req requests.PropertyRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	req.UtilityIDs = c.PostFormArray("utilities")

	form, err := c.MultipartForm()
	if err != nil && err != http.ErrNotMultipart {
		utils.Error(c, http.StatusBadRequest, "Invalid media", nil)
		return
	}

	var imageFiles []*multipart.FileHeader
	var videoFiles []*multipart.FileHeader
	if form != nil {
		imageFiles = form.File["images"]
		videoFiles = form.File["videos"]
	}

	if err := h.propertyService.Create(
		c.GetString("user_id"),
		c.GetString("role"),
		req,
		imageFiles,
		videoFiles,
	); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Created successfully", nil)
}

// GetByID handles property detail retrieval for admin and landlord roles.
func (h *PropertyHandler) GetByID(c *gin.Context) {
	propertyID := c.Param("id")

	data, err := h.propertyService.GetByID(
		c.GetString("user_id"),
		c.GetString("role"),
		propertyID,
	)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Get property successfully", data)
}

// Update handles property updates for admin and landlord roles.
func (h *PropertyHandler) Update(c *gin.Context) {
	propertyID := c.Param("id")

	var req requests.PropertyRequest
	if err := c.ShouldBind(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	req.UtilityIDs = c.PostFormArray("utilities")
	req.RemovedImageIDs = c.PostFormArray("removed_image_ids")
	req.RemovedVideoIDs = c.PostFormArray("removed_video_ids")
	if req.UtilityIDs == nil {
		req.UtilityIDs = []string{}
	}

	form, err := c.MultipartForm()
	if err != nil && err != http.ErrNotMultipart {
		utils.Error(c, http.StatusBadRequest, "Invalid media", nil)
		return
	}

	var imageFiles []*multipart.FileHeader
	var videoFiles []*multipart.FileHeader
	if form != nil {
		imageFiles = form.File["images"]
		videoFiles = form.File["videos"]
	}

	if err := h.propertyService.Update(
		c.GetString("user_id"),
		c.GetString("role"),
		propertyID,
		req,
		imageFiles,
		videoFiles,
	); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Updated successfully", nil)
}

// Delete handles property deletion and media cleanup.
func (h *PropertyHandler) Delete(c *gin.Context) {
	propertyID := c.Param("id")

	if err := h.propertyService.Delete(
		c.GetString("user_id"),
		c.GetString("role"),
		propertyID,
	); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Deleted successfully", nil)
}

// Search handles property search with pagination, filters, and sorting.
func (h *PropertyHandler) Search(c *gin.Context) {
	var req requests.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	data, err := h.propertyService.Search(
		c.GetString("user_id"),
		c.GetString("role"),
		req,
	)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Search property successfully", data)
}
