package main

import (
	"net/http"

	"wallet/internals/core/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", HomeHandler)
	router.GET("/check-saldos", services.CheckSaldos)
	err := router.Run()
	if err != nil {
		return
	}
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Wallet Balance API!"})
}
