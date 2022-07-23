package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
)

func FarmerRoute(route *gin.Engine) {
	farmer := route.Group("api/farmer")
	{
		farmer.POST("/register", controllers.RegisterFarmer)
		farmer.GET("/all", controllers.FindAllFarmer)
		farmer.GET("/:email", controllers.FindByFarmerEmail)

	}

}
