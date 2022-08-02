package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/middleware"
)

func CustomerRoute(route *gin.Engine) {
	customer := route.Group("api/customer")
	{
		customer.POST("/login", controllers.GenerateCustomerToken)
		customer.POST("/register", controllers.RegisterCustomer)
		customer.GET("/all", middleware.Auth(), controllers.FindAllCustomer)
		customer.GET("/:email", middleware.Auth(), controllers.FindByCustomerEmail)
		customer.PUT("/:email", middleware.Auth(), controllers.UpdateCustomer)
		customer.DELETE("/:email", middleware.Auth(), controllers.DeleteCustomer)
	}
}
