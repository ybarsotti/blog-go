package comment_handler

import "github.com/gin-gonic/gin"

type CommentRestHandler interface {
	PostComment(c *gin.Context)
	GetComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}
