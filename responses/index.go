package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "",
		"data":    data,
	})
	ctx.Abort()
}

func ERROR(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": err.Error(),
	})
	ctx.Abort()
}
