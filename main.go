package main

import (
	"SIAPEL/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This is home",
			"status":  "success",
		})

	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
