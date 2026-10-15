package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(gin.Recovery())

	users := map[int]User{
		1: {ID: 1, Name: "KimYoonsik"},
		2: {ID: 2, Name: "ParkSiwoo"},
	}

	r.GET("/users/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid user id",
			})
			return
		}

		user, exists := users[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	return r
}
