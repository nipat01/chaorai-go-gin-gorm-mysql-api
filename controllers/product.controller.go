package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/models"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/responses"
)

func CreateProduct(ctx *gin.Context) {
	product := models.Product{}

	if err := ctx.ShouldBindJSON(&product); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	record := database.Instance.Create(&product)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, product)

}

func FindAllProduct(ctx *gin.Context) {
	product := []models.Product{}
	var err error

	if err = database.Instance.Find(&product).Error; err != nil {
		responses.ERROR(ctx, err)
		return
	}

	responses.JSON(ctx, product)

}

func FindProductByProductId(ctx *gin.Context) {
	log.Println("FindProductByProductId: ")
	id := ctx.Param("id")
	log.Println("farmerId: ", id)
	product := models.Product{}

	record := database.Instance.Where("id = ?", id).First(&product)

	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	log.Println("product: ", product)
	responses.JSON(ctx, product)
}

func FindProductByFarmerId(ctx *gin.Context) {
	log.Println("FindProductByFarmerId: ")
	farmerId := ctx.Param("farmerId")
	log.Println("farmerId: ", farmerId)
	product := []models.Product{}

	record := database.Instance.Where("farmer_id = ?", farmerId).Find(&product)

	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	log.Println("products[]: ", product)
	responses.JSON(ctx, product)
}

func UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	log.Println("productId: ", productId)

	product := models.Product{}

	record := database.Instance.Where("id = ?", productId).First(&product)

	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	if err := ctx.ShouldBindJSON(&product); err != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	log.Println("before product: ", product)
	updateRecord := database.Instance.Save(product)
	if updateRecord.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	log.Println("after product: ", product)
	responses.JSON(ctx, product)

}

func DeleteProduct(ctx *gin.Context) {

	productId := ctx.Param("id")
	log.Println("productId: ", productId)

	product := models.Product{}

	record := database.Instance.Where("id = ?", productId).Delete(&product)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	responses.JSON(ctx, "delete product success")
}
