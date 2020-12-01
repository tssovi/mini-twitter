package db

import (
	"log"

	"./migrations"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateDatabaseConnection Utility for creating and returning a db connection
func CreateDatabaseConnection() *gorm.DB {
	log.Println("Creating connection...")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mini_twitter_db password=postgres sslmode=disable")
	migrations.RunMigrations(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully.")
	return db
}
