package handlers

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type RentalRequestHandler struct {
	service *services.RentalRequestService
}

func NewRentalRequestHandler() *RentalRequestHandler {
	return &RentalRequestHandler{service: services.NewRentalRequestService()}
}

func (h *RentalRequestHandler) Create(c *gin.Context) {
	var req requests.RentalRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	if err := h.service.Create(c.GetString("user_id"), req.PropertyID, req.Message); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Rental request created successfully", nil)
}
