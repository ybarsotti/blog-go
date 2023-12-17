package tests

import (
	"github.com/ybarsotti/blog-test/pkg/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http/httptest"
	"os"
	"testing"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	testDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}

	mainDb := db.DB
	db.DB = testDB

	handler := mainpackage.SetupRouter()
	server = httptest.NewServer(handler)

	code := m.Run()

	err = testDB.Close()
	if err != nil {
		panic("failed to close test database")
	}

	server.Close()

	db.DB = mainDb

	os.Exit(code)
}
