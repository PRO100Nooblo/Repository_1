package config

import (
	"Project1/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type handler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	dbURL := "postgres://postgres:password@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}

func New(db *gorm.DB) handler {
	return handler{db}
}
