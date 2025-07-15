package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"` // The "-" tag prevents the password from being included in JSON responses
	Email    string `json:"email" gorm:"unique"`
}
