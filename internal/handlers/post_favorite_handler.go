package handlers

import (
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostFavoriteHandler struct{ service *services.PostFavoriteService }

// NewPostFavoriteHandler creates a new PostFavoriteHandler instance.
func NewPostFavoriteHandler() *PostFavoriteHandler {
	return &PostFavoriteHandler{service: services.NewPostFavoriteService()}
}

// Add adds a post to the user's favorites.
func (h *PostFavoriteHandler) Add(c *gin.Context) {
	if err := h.service.Add(c.GetString("user_id"), c.Param("postId")); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Added post to favorites", nil)
}

// Remove removes a post from the user's favorites.
func (h *PostFavoriteHandler) Remove(c *gin.Context) {
	if err := h.service.Remove(c.GetString("user_id"), c.Param("postId")); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Removed post from favorites", nil)
}

// IDs retrieves the IDs of the user's favorite posts.
func (h *PostFavoriteHandler) IDs(c *gin.Context) {
	data, err := h.service.IDs(c.GetString("user_id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Get favorite post IDs successfully", data)
}

// Search retrieves the user's favorite posts with pagination.
func (h *PostFavoriteHandler) Search(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "12"))
	data, err := h.service.Search(c.GetString("user_id"), page, limit)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Get favorite posts successfully", data)
}
