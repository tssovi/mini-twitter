package models

import (
	"github.com/jinzhu/gorm"
)

// Migration Model representing each migration in db
type Migration struct {
	gorm.Model
	Name string `gorm:"NOT_NULL"`
}
