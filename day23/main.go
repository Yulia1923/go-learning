package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")

	users := api.Group("/users")

	users.GET("", getUsers)
	users.GET("/:id", getUser)
	users.POST("", createUser)
	users.PUT("/:id", updateUser)
	users.DELETE("/:id", deleteUser)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
