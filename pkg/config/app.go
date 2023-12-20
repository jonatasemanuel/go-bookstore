package config

import (
	"github.com/jonatasemanuel/go-bookstore/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(sqlite.Open("book-store.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	err = d.AutoMigrate(&models.Book{})

	if err != nil {
		return
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
