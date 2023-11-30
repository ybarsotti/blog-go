package comment_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/comment"
	common_handler "github.com/ybarsotti/blog-test/handler/common"
	"net/http"
	"strconv"
)

type restHandler struct {
	commentUc comment.UseCase
}

func NewCommentHandler(commentUc comment.UseCase) CommentRestHandler {
	return &restHandler{commentUc: commentUc}
}

func (h restHandler) PostComment(c *gin.Context) {
	var postData PostCommentRequest
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: "ID must be a valid integer"})
		return
	}

	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	commentEntity, err := h.commentUc.Create(id, postData.Content, postData.Author)

	if err != nil {
		c.JSON(http.StatusBadRequest, common_handler.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, common_handler.SuccessResponse{Data: commentEntity})
}
