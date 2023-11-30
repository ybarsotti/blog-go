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

	db, err := db.OpenConn()
	if err != nil {
		log.Panicln("Failed to connect database. ", err)
	}

	post_repo := repository.NewPostRepository(db)
	comment_repo := repository.NewCommentRepository(db)

	post_usecase := post.NewUseCase(post_repo)
	comment_usecase := comment.NewUseCase(comment_repo, post_repo)

	post_handler := post_handler.NewPostHandler(post_usecase)
	comment_handler := comment_handler2.NewCommentHandler(comment_usecase)

	// Post
	router.POST("/posts", post_handler.PostPost)
	router.GET("/posts", post_handler.GetPosts)
	router.GET("/posts/:id", post_handler.GetPost)
	router.PUT("/posts/:id", post_handler.UpdatePost)
	router.DELETE("/posts/:id", post_handler.DeletePost)

	// Comment
	router.POST("/posts/:id/comments", comment_handler.PostComment)

	router.Run("localhost:8009")
}
