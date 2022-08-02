package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/middleware"
)

func FarmerRoute(route *gin.Engine) {
	farmer := route.Group("api/farmer")
	{
		farmer.POST("/login", controllers.LoginFarmer)
		farmer.POST("/register", controllers.RegisterFarmer)
		farmer.GET("/all", middleware.Auth(), controllers.FindAllFarmer)
		farmer.GET("/:email", middleware.Auth(), controllers.FindByFarmerEmail)
		farmer.PUT("/:email", middleware.Auth(), controllers.UpdateFarmer)
		farmer.DELETE("/:email", middleware.Auth(), controllers.DeleteFarmer)
	}
}
