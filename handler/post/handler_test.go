package post_handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ybarsotti/blog-test/app/use_case/post"
	"github.com/ybarsotti/blog-test/pkg/db"
	"github.com/ybarsotti/blog-test/repository"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(content)

	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
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
	handler := handlerFactory()
	requestBody := []byte(`{"title": "test title", "content": "test", "author": "test"}`)
	req := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//ctx := GetTestGinContext(w)
	//MockJsonPost(ctx, requestBody)

	//handler.PostPost(ctx)

	r := gin.Default()
	r.POST("/posts", handler.PostPost)
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
