package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "users",
		})
	})

	router.Run()
}
