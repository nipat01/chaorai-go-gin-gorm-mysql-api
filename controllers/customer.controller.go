package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/models"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/responses"
)

func RegisterCustomer(ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	if err := customer.Hashpassword(customer.Password); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	record := database.Instance.Create(&customer)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, customer)
}

func FindAllCustomer(ctx *gin.Context) {
	customer := []models.Customer{}
	var err error

	if err = database.Instance.Find(&customer).Error; err != nil {
		responses.ERROR(ctx, err)
		return
	}

	responses.JSON(ctx, customer)
}

func FindByCustomerEmail(ctx *gin.Context) {
	log.Println("FindByCustomerEmail()")
	email := ctx.Param("email")
	log.Println("email: ", email)
	customer := models.Customer{}

	record := database.Instance.Where("email = ?", email).First(&customer)
	log.Println("record: ", record)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, customer)
}

func UpdateCustomer(ctx *gin.Context) {
	log.Println("UpdateCustomer(): [start] ===>")
	email := ctx.Param("email")
	log.Println("email: ", email)

	customer := models.Customer{}

	record := database.Instance.Where("email = ?", email).First(&customer)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	log.Println("customer: before =>", customer)

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	record2 := database.Instance.Save(&customer)
	if record2.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, customer)
}

func DeleteCustomer(ctx *gin.Context) {
	log.Println("UpdateCustomer(): [start] ===>")
	email := ctx.Param("email")
	customer := models.Customer{}

	record := database.Instance.Where("email = ?", email).Delete(&customer)
	log.Println("record: ", record)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, "delete success")

}
