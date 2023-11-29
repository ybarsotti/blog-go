package main

import (
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
	post_repo := repository.New(db)
	if err != nil {
		log.Panicln("Failed to connect database. ", err)
	}

	post_usecase := post.NewUseCase(post_repo)

	post_handler := post_handler.NewHandler(post_usecase)

	router.POST("/posts", post_handler.PostPost)
	router.GET("/posts", post_handler.GetPosts)
	router.GET("/posts/:id", post_handler.GetPost)
	router.PUT("/posts/:id", post_handler.UpdatePost)
	router.DELETE("/posts/:id", post_handler.DeletePost)
	router.Run("localhost:8009")
}
