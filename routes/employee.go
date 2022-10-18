package routes

import (
	"SIAPEL/config"
	"SIAPEL/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployee(ctx *gin.Context){
	var employee []models.Employee

err := ctx.ShouldBindJSON(&employee);err != nil{
	ctx.JSON(http.StatusBadRequest,gin.H{
		"status":"failed to get employee",
		"message":"internal server error",
	})
}
config.DB.Find(&employee)
}