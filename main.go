package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	fmt.Println("Serving application at http://localhost:8080")

	router.Run(":8080")
}
