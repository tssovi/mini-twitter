package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Follower Model representing a follower and following user
type Follower struct {
	gorm.Model
	UserID     uint
	FollowerID uint
}

// FollowersMigration Migration object
var FollowersMigration = Migration{
	Name: "FollowersMigration",
	Up: func(db *gorm.DB) error {
		db.CreateTable(&Follower{})
		db.Model(&Follower{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&Follower{}).AddForeignKey("follower_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&Follower{}).AddUniqueIndex("idx_follower_id_user_id", "follower_id", "user_id")
		return nil
	},
	Down: func(db *gorm.DB) error {
		db.DropTable(&Follower{})
		return nil
	},
}
