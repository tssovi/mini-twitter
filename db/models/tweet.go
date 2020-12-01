package models

import (
	"github.com/jinzhu/gorm"
)

// Tweet Model representing a post in db
type Tweet struct {
	gorm.Model
	Content string `gorm:"NOT_NULL"`
	UserID  uint   `gorm:"NOT_NULL"`
	User    User   `gorm:"foreignkey:UserID"`
}
