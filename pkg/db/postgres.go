package db

import (
	"github.com/ybarsotti/blog-test/repository"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func OpenConn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5433 sslmode=disable"
	var err error
	DB, err = gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})
	DB = DB.Debug()
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&repository.Post{}, &repository.Comment{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return DB, nil
}
