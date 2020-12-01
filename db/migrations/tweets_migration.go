package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Tweet Model representing a post in db
type Tweet struct {
	gorm.Model
	Content string `gorm:"NOT_NULL"`
	UserID  uint   `gorm:"NOT_NULL"`
}

// TweetsMigration Migration object
var TweetsMigration = Migration{
	Name: "TweetsMigration",
	Up: func(db *gorm.DB) error {
		db.CreateTable(&Tweet{})
		db.Model(&Tweet{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		return nil
	},
	Down: func(db *gorm.DB) error {
		db.DropTable(&Follower{})
		return nil
	},
}
