package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
)

func CustomerRoute(route *gin.Engine) {
	customer := route.Group("api/customer")
	{
		customer.POST("/login", controllers.GenerateToken)
		customer.POST("/register", controllers.RegisterCustomer)
		customer.GET("/all", controllers.FindAllCustomer)
		customer.GET("/:email", controllers.FindByCustomerEmail)
		customer.PUT("/:email", controllers.UpdateCustomer)
		customer.DELETE("/:email", controllers.DeleteCustomer)

	}

}
