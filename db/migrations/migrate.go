package migrations

import (
	"log"

	"../models"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jinzhu/gorm"
)

// RunMigrations Utility method for checking for checking while migrations have been executed.
func RunMigrations(db *gorm.DB) (nil, err error) {
	if !db.HasTable(&models.Migration{}) { // Create table
		db.CreateTable(&models.Migration{})
	}
	dbMigrations := []models.Migration{}
	db.Find(&dbMigrations)
	for _, migration := range AllMigrations {
		found := false
		for _, dbMigration := range dbMigrations {
			if dbMigration.Name == migration.Name {
				found = true
				break
			}
		}
		if !found {
			log.Printf("Migrating %s\n", migration.Name)
			record := models.Migration{
				Name: migration.Name,
			}
			migration.Up(db)
			db.Create(&record)
			log.Printf("Migrated %s\n", migration.Name)
		}
	}
	return nil, nil
}
