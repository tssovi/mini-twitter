package models

import (
	"github.com/jinzhu/gorm"
)

// User Model representing
type User struct {
	gorm.Model
	//ID			   uint		  `gorm:"primaryKey"`
	Username       string     `gorm:"NOT_NULL; UNIQUE"`
	Password       string     `gorm:"NOT_NULL"`
	Name           string     `gorm:"NOT_NULL"`
	IsActive       bool       `gorm:"NOT_NULL; DEFAULT:true"`
	Followers      []Follower `gorm:"foreignkey:UserID"`
	FollowingUsers []Follower `gorm:"foreignkey:FollowerID"`
	Tweets         []Tweet    `gorm:"foreignkey:UserID"`
}
