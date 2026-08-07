package handlers

import (
	"strconv"

	"easyrent-server/internal/dto/requests"
	"easyrent-server/internal/services"
	"easyrent-server/internal/utils"

	"github.com/gin-gonic/gin"
)

type PostInteractionHandler struct {
	service *services.PostInteractionService
}

func NewPostInteractionHandler() *PostInteractionHandler {
	return &PostInteractionHandler{service: services.NewPostInteractionService()}
}

func (h *PostInteractionHandler) Social(c *gin.Context) {
	data, err := h.service.Social(c.GetString("user_id"), c.Param("id"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Get post interactions successfully", data)
}
func (h *PostInteractionHandler) Like(c *gin.Context) {
	if err := h.service.Like(c.GetString("user_id"), c.Param("id")); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Liked post successfully", nil)
}
func (h *PostInteractionHandler) Unlike(c *gin.Context) {
	if err := h.service.Unlike(c.GetString("user_id"), c.Param("id")); err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Unliked post successfully", nil)
}
func (h *PostInteractionHandler) Comments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	data, err := h.service.Comments(c.Param("id"), page, limit)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Get comments successfully", data)
}
func (h *PostInteractionHandler) Comment(c *gin.Context) {
	var req requests.PostCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}
	data, err := h.service.Comment(c.GetString("user_id"), c.Param("id"), req.Content, req.ParentCommentID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	utils.Success(c, "Commented successfully", data)
}
