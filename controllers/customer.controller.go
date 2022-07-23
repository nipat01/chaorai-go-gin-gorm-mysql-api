package controllers

import (
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
