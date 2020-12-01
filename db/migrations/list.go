package migrations

import (
	"github.com/jinzhu/gorm"
)

// Migration Type representing a single migration in the system
type Migration struct {
	Name string
	Up   func(db *gorm.DB) error
	Down func(db *gorm.DB) error
}

// AllMigrations A list maintaining all migrations in the system
var AllMigrations = []Migration{
	UsersMigration,
	FollowersMigration,
	TweetsMigration,
}
