package db

import (
	"github.com/thomas-gusewelle/go-auth/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
	var err error

	Database, err = gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})
	utils.CheckError(err)

	Database.AutoMigrate(&User{})

	return nil
}
