package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "github.com/nipat01/chaorai-go-gin-gorm-mysql-api/controllers"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/middleware"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/routes"
)

func main() {
	server := gin.Default()

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

	test := server.Group("api/test").Use(middleware.Auth())
	{
		test.GET("/", testMiddleware).Use(middleware.Auth())
	}

	routes.CustomerRoute(server)
	server.Run()
}

func testMiddleware(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "ping pong",
	})
}
