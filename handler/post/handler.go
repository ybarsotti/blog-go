package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	common_handler "github.com/ybarsotti/blog-test/handler/common"
)

type restHandler struct {
	post_uc post.UseCase
}

func NewHandler(post_uc post.UseCase) PostRestHandler {
	return &restHandler{post_uc: post_uc}
}

func (h *restHandler) PostPost(c *gin.Context) {
	var postData PostPostRequest
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.post_uc.Create(postData.Title, postData.Content, postData.Author)

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusCreated, common_handler.SuccessResponse{Data: post})
}

func (h *restHandler) GetPosts(c *gin.Context) {
	posts, err := h.post_uc.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, common_handler.ErrorResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, common_handler.SuccessResponse{Data: posts})
}

func (h *restHandler) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: "ID must be a valid integer"})
		return
	}
	post := h.post_uc.Get(id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, common_handler.ErrorResponse{Message: "Post not found"})
		return
	}

	c.JSON(http.StatusOK, common_handler.SuccessResponse{Data: post})
}

func (h *restHandler) UpdatePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var postData UpdatePostRequest
	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: "ID must be a valid integer"})
		return
	}

	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.post_uc.Update(id, postData.Title, postData.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, common_handler.SuccessResponse{Data: post})
}

func (h *restHandler) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: "ID must be a valid integer"})
		return
	}

	h.post_uc.Delete(id)

	c.Status(http.StatusNoContent)
}
