package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (customer *Customer) Hashpassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	customer.Password = string(bytes)
	return nil
}

func (customer *Customer) CheckPassword(providePassword string) error {
	log.Println("CheckPassword() [start] ===>")
	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(providePassword))
	if err != nil {
		return err
	}

	return nil
}
