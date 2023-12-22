package comment_handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/comment"
	"github.com/ybarsotti/blog-test/entity"
	"github.com/ybarsotti/blog-test/pkg/db"
	"github.com/ybarsotti/blog-test/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerFactory() CommentRestHandler {
	dbConn, err := db.OpenTestConn()
	if err != nil {
		log.Fatalln(err)
	}
	postRepo := repository.NewPostRepository(dbConn)
	commentRepo := repository.NewCommentRepository(dbConn)
	commentUseCase := comment.NewUseCase(commentRepo, postRepo)
	return NewCommentHandler(commentUseCase)
}

func createTestPost() {
	dbConn, err := db.OpenTestConn()
	if err != nil {
		log.Fatalln(err)
	}
	postRepo := repository.NewPostRepository(dbConn)
	postRepo.Create(&entity.Post{
		ID:      1,
		Title:   "Test post",
		Content: "Test content",
		Author:  "Test author",
	})
}

func setup(router *gin.Engine) *gin.Engine {
	commentRestHandler := handlerFactory()
	createTestPost()
	router.POST("/posts/:id/comments", commentRestHandler.PostComment)
	router.GET("/posts/:id/comments", commentRestHandler.GetComment)
	router.DELETE("/posts/:id/comments/:comment_id", commentRestHandler.DeleteComment)
	return router
}

type CommentResponseObject struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
	Author  string `json:"author"`
}

type PostCommentResponse struct {
	Data CommentResponseObject `json:"data"`
}

func TestPostComment(t *testing.T) {
	r := gin.Default()
	setup(r)
	requestBody := []byte(`{"content": "test content", "author": "test author"}`)
	req := httptest.NewRequest("POST", "/posts/1/comments", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseBody PostCommentResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d, but got %d", http.StatusCreated, w.Code)
	}

	if responseBody.Data.Comment != "test content" {
		t.Errorf("Expected comment %s, but got %s", "test content", responseBody.Data.Comment)
	}

	if responseBody.Data.Author != "test author" {
		t.Errorf("Expected author %s, but got %s", "test author", responseBody.Data.Author)
	}
}

type GetCommentsResponse struct {
	Data []CommentResponseObject `json:"data"`
}

func TestGetComment(t *testing.T) {
	r := gin.Default()
	setup(r)
	// Create comment
	requestBody := []byte(`{"content": "test content", "author": "test author"}`)
	req := httptest.NewRequest("POST", "/posts/1/comments", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Get comment
	req = httptest.NewRequest("GET", "/posts/1/comments", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseBody GetCommentsResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected %d, but got %d", http.StatusOK, w.Code)
	}
	if len(responseBody.Data) != 1 {
		t.Errorf("Expected comment length %d, but got %d", 1, len(responseBody.Data))
	}

	if responseBody.Data[0].Comment != "test content" {
		t.Errorf("Expected comment %s, but got %s", "test content", responseBody.Data[0].Comment)
	}

	if responseBody.Data[0].Author != "test author" {
		t.Errorf("Expected author %s, but got %s", "test author", responseBody.Data[0].Author)
	}
}

func TestDeleteComment(t *testing.T) {
	r := gin.Default()
	setup(r)
	// Create comment
	requestBody := []byte(`{"content": "test content", "author": "test author"}`)
	req := httptest.NewRequest("POST", "/posts/1/comments", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Delete comment
	req = httptest.NewRequest("DELETE", "/posts/1/comments/1", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Get comments
	reqGet := httptest.NewRequest("GET", "/posts/1/comments", nil)
	wGet := httptest.NewRecorder()
	r.ServeHTTP(wGet, reqGet)

	var responseBody GetCommentsResponse
	err := json.Unmarshal(wGet.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected %d, but got %d", http.StatusNoContent, w.Code)
	}

	if len(responseBody.Data) != 0 {
		t.Errorf("Expected length %d, but got %d", 0, len(responseBody.Data))
	}
}
