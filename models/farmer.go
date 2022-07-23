package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Farmer struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (farmer *Farmer) HashpasswordFarmer(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	farmer.Password = string(bytes)
	return nil
}

func (farmer *Farmer) CheckPasswordFarmer(providePassword string) error {
	log.Println("CheckPassword() [start] ===>")
	err := bcrypt.CompareHashAndPassword([]byte(farmer.Password), []byte(providePassword))
	if err != nil {
		return err
	}

	return nil
}