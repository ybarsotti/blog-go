package routes

import (
	"github.com/gin-gonic/gin"
	posthandler "github.com/ybarsotti/blog-test/handler/post"
)

func SetupPost(router *gin.Engine, factory func() posthandler.PostRestHandler) *gin.Engine {
	postRestHandler := factory()
	router.POST("/posts", postRestHandler.PostPost)
	router.GET("/posts", postRestHandler.GetPosts)
	router.GET("/posts/:id", postRestHandler.GetPost)
	router.PUT("/posts/:id", postRestHandler.UpdatePost)
	router.DELETE("/posts/:id", postRestHandler.DeletePost)
	return router
}
