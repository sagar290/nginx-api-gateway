package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "auth-service is running",
		})
	})

	r.GET("/validate", func(c *gin.Context) {
		// get header
		token := c.GetHeader("Authorization")

		// validate token
		if token != "Bearer token" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "authorized",
		})
	})

	err := r.Run("0.0.0.0:9090")
	if err != nil {
		return
	}
}
