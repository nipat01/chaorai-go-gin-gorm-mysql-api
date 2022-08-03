package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/routes"
)

func main() {
	server := gin.Default()
	server.Use(CORSMiddleware())
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Println("Error to load .env file")
	}
	log.Println("envs: ", envs)

	dbConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", envs["DB_USER"], envs["DB_PASS"], envs["DB_URL"], envs["DB_PORT"], envs["DB_NAME"])
	log.Println("dbConnect: ", dbConnect)
	database.Connect(dbConnect)
	database.Migrate()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "hello world",
		})
	})

	routes.CustomerRoute(server)
	routes.FarmerRoute(server)
	routes.OrderRoute(server)
	routes.ProductRoute(server)

	server.Run(":" + envs["PORT"])
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
