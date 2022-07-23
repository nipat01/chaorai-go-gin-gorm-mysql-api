package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/models"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/responses"
)

func RegisterFarmer(ctx *gin.Context) {
	var farmer models.Farmer
	if err := ctx.ShouldBindJSON(&farmer); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	if err := farmer.HashpasswordFarmer(farmer.Password); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	record := database.Instance.Create(&farmer)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, farmer)
}
func FindAllFarmer(ctx *gin.Context) {
	farmer := []models.Farmer{}
	var err error

	if err = database.Instance.Find(&farmer).Error; err != nil {
		responses.ERROR(ctx, err)
		return
	}

	responses.JSON(ctx, farmer)
}

func FindByFarmerEmail(ctx *gin.Context) {
	log.Println("FindByFarmerEmail()")
	email := ctx.Param("email")
	log.Println("email: ", email)
	farmer := models.Farmer{}

	record := database.Instance.Where("email = ?", email).First(&farmer)
	log.Println("record: ", record)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, farmer)
}
