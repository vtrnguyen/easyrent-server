package handlers

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new instance of UserHandler with the necessary dependencies.
func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(),
	}
}

// GetMe retrieves the authenticated user's information based on the user ID stored in the context.
func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetString("user_id")

	data, err := h.userService.GetMe(userID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(
		c,
		"Get me successfully",
		data,
	)
}

// UpdateMe updates the authenticated user's information based on the user ID stored in the context and the provided request data.
func (h *UserHandler) UpdateMe(c *gin.Context) {
	userID := c.GetString("user_id")

	var req requests.UpdateMeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(
			c,
			err,
		)
		return
	}

	err := h.userService.UpdateMe(userID, req)
	if err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	utils.Success(
		c,
		"Updated successfully",
		nil,
	)
}

// Search handles the search functionality for users based on the provided search criteria in the request body. It returns a list of users matching the search criteria.
func (h *UserHandler) Search(c *gin.Context) {
	var req requests.SearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	data, err := h.userService.Search(req)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(
		c,
		"Search user successfully",
		data,
	)
}
