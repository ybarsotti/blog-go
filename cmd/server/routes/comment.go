package routes

import (
	"github.com/gin-gonic/gin"
	comment_handler2 "github.com/ybarsotti/blog-test/handler/comment"
)

func SetupComment(router *gin.Engine, factory func() comment_handler2.CommentRestHandler) *gin.Engine {
	commentRestHandler := factory()
	router.POST("/posts/:id/comments", commentRestHandler.PostComment)
	router.GET("/posts/:id/comments", commentRestHandler.GetComment)
	router.DELETE("/posts/:id/comments/:comment_id", commentRestHandler.DeleteComment)
	return router
}
