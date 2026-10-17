package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/profile", JWTAuth(), func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "user ID not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "profile endpoint",
			"userID":  userID,
		})
	})

	r.Run(":8080")
}
