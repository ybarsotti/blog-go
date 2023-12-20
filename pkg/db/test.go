package db

import (
	"github.com/ybarsotti/blog-test/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func OpenTestConn() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Migrator().DropTable(&repository.Post{}, &repository.Comment{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&repository.Post{}, &repository.Comment{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
