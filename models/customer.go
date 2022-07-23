package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
