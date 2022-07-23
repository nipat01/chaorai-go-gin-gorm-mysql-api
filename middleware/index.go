package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/auth"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Auth() [start]")
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": "request does not cain an access token",
			})
			ctx.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{
				"status":  401,
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
