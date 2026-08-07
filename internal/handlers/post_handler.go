package handlers

import (
	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type PostHandler struct{ postService *services.PostService }

// NewPostHandler creates a new PostHandler instance.
func NewPostHandler() *PostHandler { return &PostHandler{postService: services.NewPostService()} }

// Create handles the creation of a new post.
func (h *PostHandler) Create(c *gin.Context) {
	var req requests.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	if err := h.postService.Create(c.GetString("user_id"), req); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Created successfully", nil)
}

// GetByID retrieves a post by its ID.
func (h *PostHandler) GetByID(c *gin.Context) {
	data, err := h.postService.GetByID(c.GetString("user_id"), c.GetString("role"), c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Get post successfully", data)
}

// GetByUserID retrieves posts by the user's ID.
func (h *PostHandler) Update(c *gin.Context) {
	var req requests.PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	if err := h.postService.Update(c.GetString("user_id"), c.Param("id"), req); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Updated successfully", nil)
}

// Delete handles the deletion of a post by its ID.
func (h *PostHandler) Delete(c *gin.Context) {
	if err := h.postService.Delete(c.GetString("user_id"), c.Param("id")); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Deleted successfully", nil)
}

// GetByUserID retrieves posts by the user's ID.
func (h *PostHandler) Search(c *gin.Context) {
	var req requests.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	data, err := h.postService.Search(c.GetString("user_id"), req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Search post successfully", data)
}
