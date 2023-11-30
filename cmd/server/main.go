package main

import (
	"github.com/ybarsotti/blog-test/app/use_case/comment"
	comment_handler2 "github.com/ybarsotti/blog-test/handler/comment"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	post_handler "github.com/ybarsotti/blog-test/handler/post"
	"github.com/ybarsotti/blog-test/pkg/db"
	"github.com/ybarsotti/blog-test/repository"
)

func main() {
	router := gin.Default()

	dbConn, err := db.OpenConn()
	if err != nil {
		log.Panicln("Failed to connect database. ", err)
	}

	postRepo := repository.NewPostRepository(dbConn)
	commentRepo := repository.NewCommentRepository(dbConn)

	postUsecase := post.NewUseCase(postRepo)
	commentUsecase := comment.NewUseCase(commentRepo, postRepo)

	postRestHandler := post_handler.NewPostHandler(postUsecase)
	commentHandler := comment_handler2.NewCommentHandler(commentUsecase)

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

	err = router.Run("localhost:8009")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
