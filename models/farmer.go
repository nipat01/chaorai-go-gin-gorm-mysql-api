package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Farmer struct {
	gorm.Model
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Province    string `json:"province"`
	District    string `json:"district"`
	SubDistrict string `json:"subDistrict"`
	Postcode    string `json:"postcode"`
	PhoneNo     string `json:"phoneNo"`
	Status      string `json:"status"`
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
