package handlers

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type UtilityHandler struct {
	utilityService *services.UtilityService
}

// NewUtilityHandler creates a new instance of UtilityHandler with the necessary dependencies.
func NewUtilityHandler() *UtilityHandler {
	return &UtilityHandler{
		utilityService: services.NewUtilityService(),
	}
}

// GetAll retrieves all utilities and returns them in the response.
func (h *UtilityHandler) GetAll(c *gin.Context) {
	data, err := h.utilityService.GetAll()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Retrieved successfully", data)
}

// Create handles the creation of a new utility. It validates the request, calls the UtilityService to create the utility, and returns a success response if successful.
func (h *UtilityHandler) Create(c *gin.Context) {
	var utility requests.UtilityRequest

	if err := c.ShouldBindJSON(&utility); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	if err := h.utilityService.Create(&utility); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Created successfully", utility)
}

// Update updates an existing utility based on the provided request data. It validates the request, calls the UtilityService to update the utility, and returns a success response if successful.
func (h *UtilityHandler) Update(c *gin.Context) {
	var utility requests.UtilityRequest

	if err := c.ShouldBindJSON(&utility); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	if err := h.utilityService.Update(&utility); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Updated successfully", utility)
}

// Delete removes a utility based on the provided utility ID. It calls the UtilityService to delete the utility and returns a success response if successful.
func (h *UtilityHandler) Delete(c *gin.Context) {
	utilityID := c.Param("id")

	if err := h.utilityService.Delete(utilityID); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, "Deleted successfully", nil)
}
