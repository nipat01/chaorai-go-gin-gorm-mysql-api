package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/models"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/responses"
)

func CreateOrder(ctx *gin.Context) {
	OrderRequest := models.OrderRequest{}

	if err := ctx.ShouldBindJSON(&OrderRequest); err != nil {
		responses.ERROR(ctx, err)
		return
	}
	log.Println("order: ", OrderRequest)

	order := models.Order{}
	order.FarmerId = OrderRequest.FarmerId
	order.CustomerId = OrderRequest.CustomerId
	order.Status = OrderRequest.Status
	order.Chanel = OrderRequest.Chanel

	orderRecord := database.Instance.Create(&order)
	if orderRecord.Error != nil {
		responses.ERROR(ctx, orderRecord.Error)
		return
	}

	log.Println("orderRecord: ", orderRecord)
	for index, orderListElement := range OrderRequest.OrderList {
		log.Println("index: ", index, "orderListElement: ", orderListElement)
		orderList := models.OrderList{}
		orderList.OrderId = order.ID
		orderList.ProductId = orderListElement.ProductId
		orderList.ProductName = orderListElement.ProductName
		orderList.Price = orderListElement.Price
		orderList.Qty = orderListElement.Qty

		orderListRecord := database.Instance.Create(&orderList)
		if orderListRecord.Error != nil {
			responses.ERROR(ctx, orderRecord.Error)
			return
		}
		OrderRequest.OrderList[index].ID = orderList.ID
	}

	OrderRequest.ID = order.ID

	responses.JSON(ctx, OrderRequest)
}

func FindAllOrder(ctx *gin.Context) {
	log.Println("FindAllOrder: ")
	orderResponse := []models.OrderResponse{}
	order := []models.Order{}

	orderRecord := database.Instance.Find(&order)
	log.Println("order: ", order)
	if orderRecord.Error != nil {
		responses.ERROR(ctx, orderRecord.Error)
		return
	}

	for _, element := range order {
		log.Println("element: ", element)
		orderObject := models.OrderResponse{}
		orderObject.ID = element.ID
		orderObject.FarmerId = element.FarmerId
		orderObject.CustomerId = element.CustomerId
		orderObject.Status = element.Status
		orderObject.Chanel = element.Chanel

		orderList := []models.OrderList{}
		orderListRecord := database.Instance.Where("id = ?", element.ID).Find(&orderList)

		if orderListRecord.Error != nil {
			responses.ERROR(ctx, orderListRecord.Error)
			return
		}

		orderObject.OrderList = orderList
		orderResponse = append(orderResponse, orderObject)
	}

	log.Println("FindAllOrder: End")
	responses.JSON(ctx, orderResponse)
}

func FindOrderByFarmerId(ctx gin.Context) {

}

func UpdateOrderByCustomerId(ctx gin.Context) {

}

func UpdateOrder(ctx gin.Context) {

}

func DeleteOrder(ctx gin.Context) {

}
