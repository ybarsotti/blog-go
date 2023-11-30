package post_handler

import "github.com/gin-gonic/gin"

type PostRestHandler interface {
	PostPost(c *gin.Context)
	GetPosts(c *gin.Context)
	GetPost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}
