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
		orderListRecord := database.Instance.Where("order_id = ?", element.ID).Find(&orderList)

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

func FindOrderByOrderId(ctx *gin.Context) {
	log.Println("FindOrderByOrderId: ")
	id := ctx.Param("id")
	log.Println("id: ", id)

	orderResponse := []models.OrderResponse{}
	order := []models.Order{}
	orderRecord := database.Instance.Where("id = ?", id).Find(&order)
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
		orderListRecord := database.Instance.Where("order_id = ?", element.ID).Find(&orderList)

		if orderListRecord.Error != nil {
			responses.ERROR(ctx, orderListRecord.Error)
			return
		}

		orderObject.OrderList = orderList
		orderResponse = append(orderResponse, orderObject)
	}

	log.Println("FindOrderByOrderId: End")
	responses.JSON(ctx, orderResponse)
}

func FindOrderByFarmerEmail(ctx *gin.Context) {
	log.Println("FindOrderByFarmerEmail: ")
	farmerEmail := ctx.Param("farmerEmail")
	log.Println("farmerEmail: ", farmerEmail)

	orderResponse := []models.OrderResponse{}
	order := []models.Order{}
	orderRecord := database.Instance.Where("farmer_id = ?", farmerEmail).Find(&order)
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
		orderListRecord := database.Instance.Where("order_id = ?", element.ID).Find(&orderList)

		if orderListRecord.Error != nil {
			responses.ERROR(ctx, orderListRecord.Error)
			return
		}

		orderObject.OrderList = orderList
		orderResponse = append(orderResponse, orderObject)
	}

	log.Println("FindOrderByFarmerId: End")
	responses.JSON(ctx, orderResponse)
}

func FindOrderByCustomerEmail(ctx *gin.Context) {
	log.Println("FindOrderByCustomerEmail: ")
	customerEmail := ctx.Param("customerEmail")
	log.Println("customerEmail: ", customerEmail)

	orderResponse := []models.OrderResponse{}
	order := []models.Order{}
	orderRecord := database.Instance.Where("customer_id = ?", customerEmail).Find(&order)
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
		orderListRecord := database.Instance.Where("order_id = ?", element.ID).Find(&orderList)

		if orderListRecord.Error != nil {
			responses.ERROR(ctx, orderListRecord.Error)
			return
		}

		orderObject.OrderList = orderList
		orderResponse = append(orderResponse, orderObject)
	}

	log.Println("FindOrderByCustomerId: End")
	responses.JSON(ctx, orderResponse)

}

func UpdateOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	orderRequest := models.OrderRequest{}
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	//order
	orderRecord := database.Instance.Where("id = ?", orderId).First(&order)
	if orderRecord.Error != nil {
		responses.ERROR(ctx, orderRecord.Error)
		return
	}
	log.Println("order: before", order)

	order.CheckOrderValueChanged(orderRequest)

	updateOrderRecord := database.Instance.Save(order)
	if updateOrderRecord.Error != nil {
		responses.ERROR(ctx, updateOrderRecord.Error)
		return
	}
	log.Println("order: after", orderRequest)

	// ไม่ควรจะอัพเดท orderList ได้
	//orderList
	// for _, element := range orderRequest.OrderList {
	// 	orderList := models.OrderList{}
	// 	log.Println("element.ID: ", element.ID)

	// 	orderListRecord := database.Instance.Where("id = ?", element.ID).First(&orderList)
	// 	if orderListRecord.Error != nil {
	// 		responses.ERROR(ctx, orderListRecord.Error)
	// 		return
	// 	}

	// }

	responses.JSON(ctx, orderRequest)

}

func DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	order := models.Order{}
	OrderList := models.OrderList{}

	orderListRecord := database.Instance.Where("order_id = ?", orderId).Delete(&OrderList)
	if orderListRecord.Error != nil {
		responses.ERROR(ctx, orderListRecord.Error)
		return
	}

	orderRecord := database.Instance.Where("id = ?", orderId).Delete(&order)
	if orderRecord.Error != nil {
		responses.ERROR(ctx, orderRecord.Error)
		return
	}

	responses.JSON(ctx, "delete order and orderList success")
}

func DeleteOrderListByOrderId(ctx *gin.Context) {}

func DeleteOrderListByOrderListId(ctx *gin.Context) {
	id := ctx.Param("id")
	OrderList := models.OrderList{}
	orderListRecord := database.Instance.Where("id = ?", id).Delete(&OrderList)
	if orderListRecord.Error != nil {
		responses.ERROR(ctx, orderListRecord.Error)
		return
	}

	responses.JSON(ctx, "delete orderList success")
}
