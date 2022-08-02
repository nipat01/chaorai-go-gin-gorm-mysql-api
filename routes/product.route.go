package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
	// "github.com/nipat01/chaorai-go-gin-gorm-mysql-api/middleware"
)

func ProductRoute(route *gin.Engine) {
	product := route.Group("api/product")
	// product := route.Group("api/product").Use(middleware.Auth())
	{
		product.GET("/all", controllers.FindAllProduct)
		product.GET("/:id", controllers.FindProductByProductId)
		product.POST("/", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
		product.GET("/farmer/:farmerId", controllers.FindProductByFarmerId)
	}
}
