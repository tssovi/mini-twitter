package models

import (
	"github.com/jinzhu/gorm"
)

// Follower Model representing a follower and following user
type Follower struct {
	gorm.Model
	UserID       uint // The user who is followed
	FollowerID   uint // The user who is following
	User         User `gorm:"foreignkey:UserID"`
	FollowerUser User `gorm:"foreignkey:FollowerID"`
}
