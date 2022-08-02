package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Lastname   string `json:"lastname"`
	Province   string `json:"province"`
	District   string `json:"district"`
	SubDistrict string `json:"subDistrict"`
	Postcode   string `json:"postcode"`
	PhoneNo    string `json:"phoneNo"`
	Status     string `json:"status"`
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
