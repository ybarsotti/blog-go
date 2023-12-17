package post_handler

import (
	"github.com/ybarsotti/blog-test/app/use_case/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	common_handler "github.com/ybarsotti/blog-test/handler/common"
)

type restHandler struct {
	postUc post.UseCase
}

func NewPostHandler(postUc post.UseCase) PostRestHandler {
	return &restHandler{postUc: postUc}
}

func (h *restHandler) PostPost(c *gin.Context) {
	var postData PostPostRequest

	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.postUc.Create(postData.Title, postData.Content, postData.Author)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, common_handler.SuccessResponse{Data: post})
}

func (h *restHandler) GetPosts(c *gin.Context) {
	posts, err := h.postUc.GetAll()

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, common_handler.SuccessResponse{Data: posts})
}

func (h *restHandler) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: "ID must be a valid integer"})
		return
	}
	post := h.postUc.Get(id)
	if post.ID == 0 {
		c.Error(&common.NotFoundError{Resource: "Post"})
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

	post, err := h.postUc.Update(id, postData.Title, postData.Content)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, common_handler.SuccessResponse{Data: post})
}

func (h *restHandler) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: "ID must be a valid integer"})
		return
	}

	h.postUc.Delete(id)

	c.Status(http.StatusNoContent)
}
