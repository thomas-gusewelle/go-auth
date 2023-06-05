package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AccessToken     string
	TokenExpiration time.Time
}
