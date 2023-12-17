package main

import (
	"github.com/ybarsotti/blog-test/app/use_case/comment"
	"github.com/ybarsotti/blog-test/app/use_case/common"
	comment_handler2 "github.com/ybarsotti/blog-test/handler/comment"
	common_handler "github.com/ybarsotti/blog-test/handler/common"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	post_handler "github.com/ybarsotti/blog-test/handler/post"
	"github.com/ybarsotti/blog-test/pkg/db"
	"github.com/ybarsotti/blog-test/repository"
)

func errorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0].Err

			switch err.(type) {
			case *common.NotFoundError:
				c.JSON(http.StatusNotFound, common_handler.ErrorResponse{Message: err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		}

		c.Abort()
	}
}

func getDBConnection() *gorm.DB {
	dbConn, err := db.OpenConn()
	if err != nil {
		log.Panicln("Failed to connect database. ", err)
	}
	return dbConn
}

func postFactory() post_handler.PostRestHandler {
	dbConn := getDBConnection()
	postRepo := repository.NewPostRepository(dbConn)
	postUseCase := post.NewUseCase(postRepo)
	return post_handler.NewPostHandler(postUseCase)
}

func commentFactory() comment_handler2.CommentRestHandler {
	dbConn := getDBConnection()
	postRepo := repository.NewPostRepository(dbConn)
	commentRepo := repository.NewCommentRepository(dbConn)
	commentUseCase := comment.NewUseCase(commentRepo, postRepo)
	return comment_handler2.NewCommentHandler(commentUseCase)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(errorHandlerMiddleware())
	postRestHandler := postFactory()
	commentHandler := commentFactory()

	// Post
	router.POST("/posts", postRestHandler.PostPost)
	router.GET("/posts", postRestHandler.GetPosts)
	router.GET("/posts/:id", postRestHandler.GetPost)
	router.PUT("/posts/:id", postRestHandler.UpdatePost)
	router.DELETE("/posts/:id", postRestHandler.DeletePost)

	// Comment
	router.POST("/posts/:id/comments", commentHandler.PostComment)
	router.GET("/posts/:id/comments", commentHandler.GetComment)
	router.DELETE("/posts/:id/comments/:comment_id", commentHandler.DeleteComment)
	return router
}

func main() {
	router := SetupRouter()
	err := router.Run("localhost:8009")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
