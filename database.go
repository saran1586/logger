package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Employee struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

func ConnectDB() {

	// Connection string comes from DATABASE_URL; the fallback is a local development database.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	DB.AutoMigrate(&Employee{})

	log.Println("Database Connected")
}
