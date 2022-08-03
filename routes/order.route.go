package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
	// "github.com/nipat01/chaorai-go-gin-gorm-mysql-api/middleware"
)

func OrderRoute(route *gin.Engine) {
	order := route.Group("/api/order")
	// order := route.Group("/api/order").Use(middleware.Auth())
	{
		order.GET("/", controllers.FindAllOrder)
		order.GET("/:id", controllers.FindOrderByOrderId)
		order.POST("/", controllers.CreateOrder)
		order.PUT("/:orderId", controllers.UpdateOrder)
		order.DELETE("/:orderId", controllers.DeleteOrder)
		order.DELETE("/list/:id", controllers.DeleteOrderListByOrderListId)
		order.GET("/customer/:customerEmail", controllers.FindOrderByCustomerEmail)
		order.GET("/farmer/:farmerEmail", controllers.FindOrderByFarmerEmail)
	}
}
