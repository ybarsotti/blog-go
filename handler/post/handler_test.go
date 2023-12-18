package post_handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	"github.com/ybarsotti/blog-test/pkg/db"
	"github.com/ybarsotti/blog-test/repository"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlerFactory() PostRestHandler {
	dbConn, err := db.OpenTestConn()
	if err != nil {
		log.Fatalln(err)
	}
	postRepo := repository.NewPostRepository(dbConn)
	postUseCase := post.NewUseCase(postRepo)
	return NewPostHandler(postUseCase)
}

func setup(router *gin.Engine) *gin.Engine {
	postRestHandler := handlerFactory()
	router.POST("/posts", postRestHandler.PostPost)
	router.GET("/posts", postRestHandler.GetPosts)
	router.GET("/posts/:id", postRestHandler.GetPost)
	router.PUT("/posts/:id", postRestHandler.UpdatePost)
	router.DELETE("/posts/:id", postRestHandler.DeletePost)
	return router
}

type PostPostResponseObject struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostPostResponse struct {
	Data PostPostResponseObject `json:"data"`
}

func TestPostPostHandler(t *testing.T) {
	r := gin.Default()
	setup(r)
	requestBody := []byte(`{"title": "test title", "content": "test", "author": "test"}`)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d, but got %d", http.StatusCreated, w.Code)
	}
	var responseBody PostPostResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}
	if responseBody.Data.ID != 1 {
		t.Errorf("Expected ID to be 1")
	}
	if responseBody.Data.Title != "test title" {
		t.Errorf("Expected title to be %s, received %s", "test title", responseBody.Data.Title)
	}
	if responseBody.Data.Author != "test" {
		t.Errorf("Expected author to be %s, received %s", "test", responseBody.Data.Author)
	}
	if responseBody.Data.Content != "test" {
		t.Errorf("Expected content to be %s, received %s", "test", responseBody.Data.Content)
	}
}

type GetPostsResponseObject struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetPostsResponse struct {
	Data []GetPostsResponseObject `json:"data"`
}

func TestGetPosts(t *testing.T) {
	r := gin.Default()
	setup(r)
	requestBody := []byte(`{"title": "test title 2", "content": "test 2", "author": "test 2"}`)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/posts", nil)
	r.ServeHTTP(w, req2)

	var responseBody GetPostsResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, received %d", w.Code)
	}

	if len(responseBody.Data) != 1 {
		t.Errorf("Expected length of %d, received %d", 1, len(responseBody.Data))
	}
}

type GetPostResponseObject struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetPostResponse struct {
	Data GetPostResponseObject `json:"data"`
}

func TestGetPost(t *testing.T) {
	r := gin.Default()
	setup(r)
	requestBody := []byte(`{"title": "test title 2", "content": "test 2", "author": "test 2"}`)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req2 := httptest.NewRequest("GET", "/posts/1", nil)
	r.ServeHTTP(w, req2)

	var responseBody GetPostResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, received %d", w.Code)
	}

	if responseBody.Data.ID != 1 {
		t.Errorf("Expected ID of %d, received %d", 1, responseBody.Data.ID)
	}

	if responseBody.Data.Title != "test title 2" {
		t.Errorf("Expected title %s, received %s", "test title 2", responseBody.Data.Title)
	}
}

type UpdatePostResponseObject struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type UpdatePostResponse struct {
	Data UpdatePostResponseObject `json:"data"`
}

func TestUpdatePost(t *testing.T) {
	r := gin.Default()
	setup(r)
	requestBody := []byte(`{"title": "test title 2", "content": "test 2", "author": "test 2"}`)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	requestPutBody := []byte(`{"title": "test update", "content": "test update", "author": "test update"}`)
	req2 := httptest.NewRequest("PUT", "/posts/1", bytes.NewBuffer(requestPutBody))
	r.ServeHTTP(w, req2)

	var responseBody UpdatePostResponse
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, received %d", w.Code)
	}

	if responseBody.Data.ID != 1 {
		t.Errorf("Expected ID of %d, received %d", 1, responseBody.Data.ID)
	}

	if responseBody.Data.Title != "test update" {
		t.Errorf("Expected title %s, received %s", "test update", responseBody.Data.Title)
	}

	if responseBody.Data.Content != "test update" {
		t.Errorf("Expected content %s, received %s", "test update", responseBody.Data.Content)
	}

	if responseBody.Data.Author != "test 2" {
		t.Errorf("Expected author %s, received %s", "test 2", responseBody.Data.Author)
	}
}

func TestDeletePost(t *testing.T) {
	r := gin.Default()
	setup(r)
	requestBody := []byte(`{"title": "test title 2", "content": "test 2", "author": "test 2"}`)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	w = httptest.NewRecorder()
	req2 := httptest.NewRequest("DELETE", "/posts/1", nil)
	r.ServeHTTP(w, req2)

	w2 := httptest.NewRecorder()
	req3 := httptest.NewRequest("GET", "/posts", nil)
	r.ServeHTTP(w2, req3)

	var responseBodyGet GetPostsResponse
	err := json.Unmarshal(w2.Body.Bytes(), &responseBodyGet)
	if err != nil {
		t.Fatal(err)
	}

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status code 204, received %d", w.Code)
	}

	if len(responseBodyGet.Data) != 0 {
		t.Errorf("Expected len of %d, received %d", 0, len(responseBodyGet.Data))
	}
}
