package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
)

func main() {
	server := gin.Default()
	
	database.Connect("root:123456@tcp(localhost:3306)/db-mysql-dev?parseTime=true")
	database.Migrate()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "hello world",
		})
	})

	server.Run()
}
