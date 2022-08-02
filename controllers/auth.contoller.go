package controllers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/auth"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/database"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/models"
	"github.com/nipat01/chaorai-go-gin-gorm-mysql-api/responses"
	"log"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateCustomerToken(ctx *gin.Context) {
	log.Println("GenerateToken() [start]")
	var request TokenRequest
	var customer models.Customer

	if err := ctx.ShouldBindJSON(&request); err != nil {
		responses.ERROR(ctx, err)
		return
	}

	log.Println("request: ", request)
	// check email exist and password is matched
	record := database.Instance.Where("email = ?", request.Email).First(&customer)
	log.Println("record: ", record)
	if record.Error != nil {
		responses.ERROR(ctx, record.Error)
		return
	}

	credentailError := customer.CheckPassword(request.Password)
	if credentailError != nil {
		responses.ERROR(ctx, credentailError)
		return
	}
	log.Println("credentailError: ", credentailError)

	tokenString, err := auth.GenerateJWT(customer.Email, customer.Name)

	if err != nil {
		responses.ERROR(ctx, err)
		return
	}
	responses.JSON(ctx, tokenString)
}
