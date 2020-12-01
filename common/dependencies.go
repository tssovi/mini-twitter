package common

import (
	"../db"
	"github.com/jinzhu/gorm"
)

// IDependencies Interface for defining dependencies
type IDependencies struct {
	DB *gorm.DB
}

// Dependencies Instance containing dependencies
var Dependencies = IDependencies{
	DB: db.CreateDatabaseConnection(),
}
