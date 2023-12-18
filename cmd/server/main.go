package main

import (
	"github.com/ybarsotti/blog-test/app/use_case/comment"
	"github.com/ybarsotti/blog-test/app/use_case/common"
	"github.com/ybarsotti/blog-test/cmd/server/routes"
	commenthandler2 "github.com/ybarsotti/blog-test/handler/comment"
	commonhandler "github.com/ybarsotti/blog-test/handler/common"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	posthandler "github.com/ybarsotti/blog-test/handler/post"
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
				c.JSON(http.StatusNotFound, commonhandler.ErrorResponse{Message: err.Error()})
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

func postFactory() posthandler.PostRestHandler {
	dbConn := getDBConnection()
	postRepo := repository.NewPostRepository(dbConn)
	postUseCase := post.NewUseCase(postRepo)
	return posthandler.NewPostHandler(postUseCase)
}

func commentFactory() commenthandler2.CommentRestHandler {
	dbConn := getDBConnection()
	postRepo := repository.NewPostRepository(dbConn)
	commentRepo := repository.NewCommentRepository(dbConn)
	commentUseCase := comment.NewUseCase(commentRepo, postRepo)
	return commenthandler2.NewCommentHandler(commentUseCase)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(errorHandlerMiddleware())
	routes.SetupPost(router, postFactory)
	routes.SetupComment(router, commentFactory)
	return router
}

func main() {
	router := SetupRouter()
	err := router.Run("localhost:8009")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
