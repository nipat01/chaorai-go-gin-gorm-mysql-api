package database

import (
	"log"

	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {

	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal("dbError: ", dbError)
		panic("Cannot conect to DB")
	}
	log.Println("Connect to Database")
}

func Migrate() {
	Instance.AutoMigrate(&models.Customer{}, &models.Farmer{}, &models.Product{}, &models.Order{}, &models.OrderList{})
	log.Println("Database Migration Complete")
}
