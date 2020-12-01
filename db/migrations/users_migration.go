package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User Model representing
type User struct {
	gorm.Model
	Username string `gorm:"NOT_NULL; UNIQUE"`
	Password string `gorm:"NOT_NULL"`
	Name     string `gorm:"NOT_NULL"`
	IsActive bool   `gorm:"NOT_NULL; DEFAULT:true"`
}

// UsersMigration Migration object
var UsersMigration = Migration{
	Name: "UsersMigration",
	Up: func(db *gorm.DB) error {
		db.CreateTable(&User{})
		return nil
	},
	Down: func(db *gorm.DB) error {
		db.DropTable(&User{})
		return nil
	},
}
