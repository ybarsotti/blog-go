package db

import (
	"github.com/ybarsotti/blog-test/repository"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConn() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5433 sslmode=disable"
	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&repository.Post{}, &repository.Comment{})

	return db, nil
}
